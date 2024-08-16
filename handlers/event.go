package handlers

import (
	"context"
	"time"

	"github.com/Dostonlv/booking-app/models"
	"github.com/gofiber/fiber/v2"
)

type eventHandler struct {
	repository models.EventRepository
}

func NewHandler(router fiber.Router, repository models.EventRepository) {
	handler := &eventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
}

func (h *eventHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    events,
	})
}

func (h *eventHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil

}

func (h *eventHandler) GetOne(ctx *fiber.Ctx) error {
	return nil

}
