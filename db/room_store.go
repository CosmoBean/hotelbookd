package db

import (
	"context"

	"github.com/CosmoBean/hotelbookd/models"
	"github.com/CosmoBean/hotelbookd/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *models.Room) (*models.Room, error)
	GetRooms(context.Context, bson.M) ([]*models.Room, error)
}

type MongoRoomStore struct {
	roomCollection *mongo.Collection

	HotelStore *MongoHotelStore
}

func NewMongoRoomStore(client *mongo.Client, database string) *MongoRoomStore {
	roomCollection := utils.GetEnvDefault("MONGO_ROOM_COLLECTION", "rooms")
	hotelCollection := utils.GetEnvDefault("MONGO_HOTELS_COLLECTION", "hotels")
	hotelStore := &MongoHotelStore{hotelCollection: client.Database(database).Collection(hotelCollection)}
	return &MongoRoomStore{
		roomCollection: client.Database(database).Collection(roomCollection),
		HotelStore:     hotelStore,
	}
}

func (r *MongoRoomStore) InsertRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	resp, err := r.roomCollection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = resp.InsertedID.(primitive.ObjectID)
	if err := r.HotelStore.AddRoom(ctx, room.HotelId, room.Id); err != nil {
		return nil, err
	}
	return room, nil
}

func (r *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*models.Room, error) {
	resp, err := r.roomCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*models.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
