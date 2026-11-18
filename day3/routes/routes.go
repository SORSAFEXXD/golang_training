package routes

import (
	"github.com/gofiber/fiber/v2"

	"day3/controllers"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	//users
	app.Get("/", controllers.HelloTest)
	v1.Get("/", controllers.StringTest)
	v1.Post("/", controllers.BodyParserTest)
	v1.Get("/user/:name", controllers.ParamsTest)
	v1.Post("/inet", controllers.QuerryTest)
	v1.Post("/valid", controllers.ValidTest)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", controllers.GetDogs)
	dog.Get("/deleted", controllers.GetDeletedDogs)
	dog.Get("/range", controllers.GetDogsByIDRange)
	dog.Get("/filter", controllers.GetDog)
	dog.Get("/json", controllers.GetDogsJson)
	dog.Post("/", controllers.AddDog)
	dog.Put("/:id", controllers.UpdateDog)
	dog.Delete("/:id", controllers.RemoveDog)

	//CRUD company
	company := v1.Group("/company")
	company.Get("", controllers.GetCompanies)
	company.Get("/:id", controllers.GetCompany)
	company.Post("/", controllers.AddCompany)
	company.Put("/:id", controllers.UpdateCompany)
	company.Delete("/:id", controllers.RemoveCompany)
}
