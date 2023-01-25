package instance

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"

	log "gorm.io/gorm/logger"
)

type gormLogger struct {
	Logger                *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	ShowErrRecordNotFound bool
}

func (g *gormLogger) LogMode(lvl log.LogLevel) log.Interface {
	if lvl <= log.Warn {
		g.Logger.SetLevel(logrus.WarnLevel)
	} else {
		g.Logger.SetLevel(logrus.TraceLevel)
	}
	return g
}

func (g *gormLogger) Info(
	ctx context.Context,
	msg string,
	args ...interface{},
) {
	g.Logger.WithContext(ctx).Infof(msg, args)
}

func (g *gormLogger) Warn(
	ctx context.Context,
	msg string,
	args ...interface{},
) {
	g.Logger.WithContext(ctx).Warnf(msg, args)
}

func (g *gormLogger) Error(
	ctx context.Context,
	msg string,
	args ...interface{},
) {
	g.Logger.WithContext(ctx).Errorf(msg, args)
}

func (g *gormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	executionTime := time.Since(begin)
	query, _ := fc()
	fields := logrus.Fields{}

	if g.SourceField != "" {
		fields[g.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && g.ShowErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		g.Logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", query, executionTime)
		return
	}

	if g.SlowThreshold > 0 && executionTime > g.SlowThreshold {
		g.Logger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", query, executionTime)
		return
	}

	g.Logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", query, executionTime)
}

func NewPostgreSQLInstance(
	logger *logrus.Logger,
	dsn string,
) (*gorm.DB, *sql.DB) {
	var level log.LogLevel
	lvl := logrus.GetLevel()

	switch lvl {
	case logrus.ErrorLevel:
		level = log.Error
	case logrus.WarnLevel:
		level = log.Warn
	case logrus.InfoLevel:
		level = log.Info
	default:
		level = log.Info
	}

	gl := gormLogger{Logger: logger}

	postgreDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gl.LogMode(level)})
	if err != nil {
		logger.Fatalf("Fatal to open database connection! Err: %s", err)
	}

	sqlDb, err := postgreDb.DB()
	if err != nil {
		logger.Fatalf("Database Error! Err: %s", err)
	}

	return postgreDb, sqlDb
}
