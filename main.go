package main

import (
	"backend/database"
	"backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env File")
	}

}
func main() {
	//
	port := os.Getenv("APP_PORT")
	// Database Connect
	database.Connect()

	// App init
	app := fiber.New()

	// Use CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Post Routing
	routes.EventRoute(app)

	// Go Fiber Listener
	app.Listen(":" + port)
}
