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
	DBInstance *mongo.Database
	once       sync.Once
)

func getInstance() *mongo.Database {
	once.Do(func() {
		var (
			mongoHost   = utils.GetEnvDefault("MONGO_HOST", "localhost")
			mongoPort   = utils.GetEnvDefault("MONGO_PORT", "27017")
			mongoDBName = utils.GetEnvDefault("MONGO_DBNAME", "hotel-reservation")
		)
		var dbUri = fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
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
