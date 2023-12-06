package db

import (
	"context"

	"github.com/CosmoBean/hotelbookd/models"
	"github.com/CosmoBean/hotelbookd/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *models.Hotel) (*models.Hotel, error)
	AddRoom(context.Context, primitive.ObjectID, primitive.ObjectID) error
	GetHotels(context.Context, bson.M) ([]*models.Hotel, error)
}

type MongoHotelStore struct {
	hotelCollection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, database string) *MongoHotelStore {
	hotelCollection := utils.GetEnvDefault("MONGO_HOTEL_COLLECTION", "hotels")
	return &MongoHotelStore{
		hotelCollection: client.Database(database).Collection(hotelCollection),
	}
}

func (h *MongoHotelStore) InsertHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	resp, err := h.hotelCollection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.Id = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (h *MongoHotelStore) AddRoom(ctx context.Context, hotelId primitive.ObjectID, roomId primitive.ObjectID) error {
	_, err := h.hotelCollection.UpdateOne(ctx, bson.M{"_id": hotelId}, bson.M{"$push": bson.M{"rooms": roomId}})
	if err != nil {
		return err
	}
	return nil
}

func (h *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*models.Hotel, error) {
	resp, err := h.hotelCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*models.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}
