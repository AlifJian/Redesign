package routes

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func EventRoute(app *fiber.App) {
	// Grouping Post Route
	router := app.Group("/event")

	// GET METHOD
	router.Get("/", controller.GetAll)

	// Search
	router.Get("/:title", controller.Search)

	// POST METHOD
	router.Post("/", controller.Create)

	// UPDATE METHOD
	router.Put("/:id", controller.Update)

	// Delete Method
	router.Delete("/:id", controller.Delete)
}
