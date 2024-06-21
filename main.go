package main

import (
	"backend/database"
	"backend/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	// Database Connect
	db := database.Connect()

	// App init
	app := fiber.New()

	// Use CSRF Middleware
	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "csrf_",               // name of the session cookie
		CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
		Expiration:     1 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUIDv4,          // creates a new CSRF token
	}

	app.Use(csrf.New(csrfConfig))

	// Post Routing
	routes.PostRoute(app, db)

	// Go Fiber Listener
	app.Listen(":8000")
}
