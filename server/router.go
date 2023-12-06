package server

import (
	"log"

	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/handler"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/gofiber/fiber/v2"
)

var errorConfig = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func Init() {
	listenAddr := utils.GetEnvDefault("API_PORT", ":8080")
	api := fiber.New(errorConfig)
	api.Get("/health", handler.HealthCheck)

	//DB Variables
	dbName := utils.GetEnvDefault("MONGO_DBNAME", "hotel-reservation")
	mongoClient := db.GetMongoClient()

	// api routes
	apiV1 := api.Group("/api/v1")

	//handlers
	userHandler := handler.NewUserHandler(db.NewMongoUserStore(mongoClient, dbName))
	hotelHandler := handler.NewHotelHandler(db.NewMongoHotelStore(mongoClient, dbName), db.NewMongoRoomStore(mongoClient, dbName))

	//usersAPI
	apiV1.Get("/users", userHandler.GetUsers)
	apiV1.Get("/users/:id", userHandler.GetUser)
	apiV1.Post("/users", userHandler.CreateNewUser)
	apiV1.Delete("/users/:id", userHandler.DeleteUser)
	apiV1.Put("/users/:id", userHandler.UpdateUser)

	//hotelAPI
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotels/:id/rooms", hotelHandler.HandleGetHotelRooms)

	err := api.Listen(listenAddr)
	if err != nil {
		log.Fatal("unable to start the fiber server, ", err)
	}

}
