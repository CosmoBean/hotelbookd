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
	users, err := h.userStore.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userStore.GetUserByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (h *UserHandler) CreateNewUser(ctx *fiber.Ctx) error {
	var createUserReq models.CreateUserRequest
	if err := ctx.BodyParser(&createUserReq); err != nil {
		return err
	}

	if err := createUserReq.Validate(); err != nil {
		return err
	}

	user, err := models.NewUserFromParams(createUserReq)
	if err != nil {
		return err
	}

	InsertedUser, err := h.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}

	return ctx.JSON(InsertedUser)
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var updateUserReq models.UpdateUserRequest
	if err := ctx.BodyParser(&updateUserReq); err != nil {
		return err
	}

	updatedUser, err := h.userStore.UpdateUser(ctx.Context(), id, updateUserReq)
	if err != nil {
		return err
	}
	return ctx.JSON(updatedUser)
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.userStore.DeleteUser(ctx.Context(), id); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
