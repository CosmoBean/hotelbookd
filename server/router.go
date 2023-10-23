package server

import (
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/handler"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Init() {
	listenAddr := utils.GetEnvDefault("API_PORT", ":8080")
	api := fiber.New()
	api.Get("/health", handler.HealthCheck)

	// api routes
	apiV1 := api.Group("/api/v1")

	userHandler := handler.NewUserHandler(db.NewMongoUserStore(db.Get()))
	apiV1.Get("/users", userHandler.GetUsers)
	apiV1.Get("/users/:id", userHandler.GetUser)

	err := api.Listen(listenAddr)
	if err != nil {
		log.Fatal("unable to start the fiber server, ", err)
	}

}
