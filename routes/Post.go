package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PostRoute(app *fiber.App, db *gorm.DB) {
	// Grouping Post Route
	router := app.Group("/post")

	// GET METHOD
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	// POST METHOD
	router.Post("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	// UPDATE METHOD
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	// Delete Method
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})
}
