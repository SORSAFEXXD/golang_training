package routes

import (
	"last_project/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	})

	user.Get("/", controllers.GetProfiles)
	user.Get("/search", controllers.SearchProfiles)
	user.Get("/age_groups", controllers.GetAgeGroups)
	user.Get("/:employee_id", controllers.GetProfile)

	user.Post("/", auth, controllers.CreateProfile)
	user.Put("/:employee_id", auth, controllers.UpdateProfile)
	user.Delete("/:employee_id", auth, controllers.DeleteProfile)
}
