package main

import (
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Database Connect
	db := database.Connect()

	// App init
	app := fiber.New()

	// Use CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Post Routing
	routes.PostRoute(app, db)

	// Go Fiber Listener
	app.Listen(":8000")
}
