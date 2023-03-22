package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haritechversant/go-sample-docker/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
	// Update an existing fact
	app.Put("/fact/:id", handlers.UpdateFact)

	// Delete an existing fact
	app.Delete("/fact/:id", handlers.DeleteFact)
}
