package main

import (
	c "day2/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	//basic auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	// --- 5.1 API V1 ---
	v1 := app.Group("/api/v1")
	v1.Get("/fact/:n", c.Calculate)

	// --- 5.2 API V3 ---
	v3 := app.Group("/api/v3")
	v3.Get("/safe" , c.ChangeWordToInt)

	// 6. API v1 method POST
	v1.Post("/register", c.Register)

	app.Listen(":3000")
}
