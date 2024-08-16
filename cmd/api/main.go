package main

import (
	"github.com/Dostonlv/booking-app/handlers"
	"github.com/Dostonlv/booking-app/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking App",
		ServerHeader: "Fiber",
	})

	// repository
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api/v1")

	// event handler
	handlers.NewHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
