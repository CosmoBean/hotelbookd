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
	GetUsers(context.Context) ([]*models.User, error)
	InsertUser(context.Context, *models.User) (*models.User, error)
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

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*models.User, error) {
	cur, err := s.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*models.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := s.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
