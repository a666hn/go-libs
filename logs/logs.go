package logs

import "github.com/sirupsen/logrus"

type defaultFieldsHook struct {
	Fields map[string]interface{}
}

func (h *defaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *defaultFieldsHook) Fire(e *logrus.Entry) error {
	for i, f := range h.Fields {
		e.Data[i] = f
	}
	return nil
}

type logConfig struct {
	IsProduction bool
	Fields       map[string]interface{}
}

type logOption func(*logConfig)

func IsProduction(isProd bool) logOption {
	return func(config *logConfig) {
		config.IsProduction = isProd
	}
}

func LogFields(fields map[string]interface{}) logOption {
	return func(config *logConfig) {
		config.Fields = fields
	}
}

func NewLoggerInstance(opts ...logOption) *logrus.Logger {
	var lvl logrus.Level
	l := logrus.New()

	lc := &logConfig{}

	for _, opt := range opts {
		opt(lc)
	}

	if lc.IsProduction {
		lvl = logrus.WarnLevel
	} else {
		lvl = logrus.TraceLevel
	}

	l.SetLevel(lvl)
	l.AddHook(&defaultFieldsHook{lc.Fields})

	return l
}
