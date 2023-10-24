package server

import (
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/handler"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/gofiber/fiber/v2"
	"log"
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

	// api routes
	apiV1 := api.Group("/api/v1")

	userHandler := handler.NewUserHandler(db.NewMongoUserStore(db.Get()))
	apiV1.Get("/users", userHandler.GetUsers)
	apiV1.Get("/users/:id", userHandler.GetUser)
	apiV1.Post("/users", userHandler.CreateNewUser)

	err := api.Listen(listenAddr)
	if err != nil {
		log.Fatal("unable to start the fiber server, ", err)
	}

}
