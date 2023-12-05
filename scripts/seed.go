package main

import (
	"context"
	"log"

	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/models"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(hotelName, hotelLocation string) {
	hotel := models.Hotel{
		Name:     hotelName,
		Location: hotelLocation,
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []models.Room{
		{
			Type:      models.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      models.DeluxeRoomType,
			BasePrice: 199.9,
		},
		{
			Type:      models.SeaSideRoomType,
			BasePrice: 129.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal("cannot insert hotel -->", err)
	}

	for _, room := range rooms {
		room.HotelId = insertedHotel.Id

		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal("cannot insert room-->", err)
		}
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("error Loading the env file: ", err)
	}
	client = db.GetMongoClient()
	dbName := utils.GetEnvDefault("MONGO_DBNAME", "hotel-reservation")
	if err := client.Database(dbName).Drop(ctx); err != nil {
		log.Fatal("unable to drop the database")
	}
	hotelStore = db.NewMongoHotelStore(client, dbName)
	roomStore = db.NewMongoRoomStore(client, dbName)
}

func main() {
	seedHotel("Bellucia", "France")
	seedHotel("the Cozy Hotel", "Netherlands")
	seedHotel("OYO", "India")
}
