package instance

// import (
// 	"context"
// 	"github.com/sirupsen/logrus"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )
//
// type mongoDb struct {
// 	ContextTimeout  context.Context
// 	MongoDbDatabase *mongo.Database
// 	MongoDbClient   *mongo.Client
// 	MongoCollection *mongo.Collection
// }
//
// type MongoDBInstance interface {
// 	GetCtx() context.Context
// 	GetDB() *mongo.Database
// 	GetCollection()
// }
//
// type mongoDbInstanceConfiguration struct {
// 	Logger *logrus.Logger
// }
//
// type MongoInstanceOption func(mg *mongoDbInstanceConfiguration)
//
// func SetLogger(logger *logrus.Logger) MongoInstanceOption {
// 	return func(mg *mongoDbInstanceConfiguration) {
// 		mg.Logger = logger
// 	}
// }
//
// func NewMongoDBInstance(
// 	mongoDsn string,
// 	dbName string,
// 	opts ...MongoInstanceOption,
// ) (*mongo.Database, *mongo.Collection) {
// 	mongoOpts := &mongoDbInstanceConfiguration{}
// 	mongoOpts.Logger = logrus.StandardLogger() // For default logger
//
// 	for _, opt := range opts {
// 		opt(mongoOpts)
// 	}
//
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDsn))
// 	if err != nil {
// 		mongoOpts.Logger.Fatalf("Fatal to connect to database %v", err)
// 	}
//
// 	database := client.Database(dbName)
// }
