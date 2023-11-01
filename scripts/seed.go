package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/models"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("error Loading the env file: ", err)
	}
	ctx := context.Background()
	client := db.GetMongoClient()
	dbName := utils.GetEnvDefault("MONGO_DBNAME", "hotel-reservation")
	hotelStore := db.NewMongoHotelStore(client, dbName)
	hotel := models.Hotel{
		Name:     "Bellucia",
		Location: "France",
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

	roomStore := db.NewMongoRoomStore(client, dbName)

	for _, room := range rooms {
		room.HotelId = insertedHotel.Id

		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal("cannot insert room-->", err)
		}
		fmt.Println(insertedRoom)
	}

	fmt.Println(insertedHotel)

}
