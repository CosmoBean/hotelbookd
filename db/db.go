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
	DBClient *mongo.Client
	err      error
	once     sync.Once
)

func getClient() *mongo.Client {
	once.Do(func() {
		var (
			mongoHost = utils.GetEnvDefault("MONGO_HOST", "localhost")
			mongoPort = utils.GetEnvDefault("MONGO_PORT", "27017")
		)
		var dbUri = fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
		DBClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
		if err != nil {
			log.Panic("error while connecting to the database ", err)
		}
	})
	return DBClient
}

func GetMongoClient() *mongo.Client {
	return getClient()
}
