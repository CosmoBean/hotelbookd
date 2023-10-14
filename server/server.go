package server

import (
	"github.com/CosmoBean/hotelbookd/handler"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Init() {
	app := fiber.New()
	app.Get("/health", handler.HealthCheck)
	err := app.Listen(":9020")
	if err != nil {
		log.Fatal("unable to start the fiber server, ", err)
	}

}
