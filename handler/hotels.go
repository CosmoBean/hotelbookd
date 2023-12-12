package handler

import (
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {
	hotels, err := h.store.Hotel.GetHotels(ctx.Context(), bson.M{})
	if err != nil {
		return err
	}
	return ctx.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotelRooms(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"hotelId": oid}
	rooms, err := h.store.Room.GetRooms(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(rooms)
}
