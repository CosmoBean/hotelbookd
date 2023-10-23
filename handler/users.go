package handler

import (
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) GetUsers(ctx *fiber.Ctx) error {
	user := models.User{
		FirstName: "James",
		LastName:  "Water",
	}
	return ctx.JSON(user)
}
func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}
