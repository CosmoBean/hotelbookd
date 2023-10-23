package handler

import (
	"github.com/CosmoBean/hotelbookd/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	user := models.User{
		FirstName: "James",
		LastName:  "Water",
	}
	return ctx.JSON(user)
}
func GetUser(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{
		"data": "user1",
	})
}
