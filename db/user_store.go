package db

import (
	"context"
	"github.com/CosmoBean/hotelbookd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*models.User, error)
}

type MongoUserStore struct {
	userCollection *mongo.Collection
}

func NewMongoUserStore(db *mongo.Database) *MongoUserStore {
	return &MongoUserStore{
		userCollection: db.Collection(UserColelction),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	var oid primitive.ObjectID
	var err error

	if oid, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	if err = s.userCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
