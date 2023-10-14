package handler

import "github.com/gofiber/fiber/v2"

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"msg": "alive"})
}
