package db

import (
	"context"
	"fmt"
	"github.com/CosmoBean/hotelbookd/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	DBInstance     *mongo.Database
	once           sync.Once
	mongoUsername  = utils.GetEnvDefault("MONGO_USERNAME", "username")
	mongoPassword  = utils.GetEnvDefault("MONGO_PASSWORD", "password")
	mongoHost      = utils.GetEnvDefault("MONGO_HOST", "localhost")
	mongoPort      = utils.GetEnvDefault("MONGO_PORT", "27017")
	mongoDBName    = utils.GetEnvDefault("MONGO_DBNAME", "hotel-reservation")
	UserColelction = utils.GetEnvDefault("MONGO_USER_COLLECTION", "users")
)

var dbUri = fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUsername, mongoPassword, mongoHost, mongoPort)

func getInstance() *mongo.Database {
	once.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
		if err != nil {
			log.Panic("error while connecting to the database ", err)
		}
		DBInstance = client.Database(mongoDBName)
	})
	return DBInstance
}

func Get() *mongo.Database {
	return getInstance()
}
