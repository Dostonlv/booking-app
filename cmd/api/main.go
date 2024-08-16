package main

import (
	"fmt"

	"github.com/Dostonlv/booking-app/config"
	"github.com/Dostonlv/booking-app/db"
	"github.com/Dostonlv/booking-app/handlers"
	"github.com/Dostonlv/booking-app/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigration)
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking App",
		ServerHeader: "Fiber",
	})

	// repository
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	// Routing
	server := app.Group("/api/v1")

	// handler
	handlers.NewHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
