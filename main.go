package main

import (
	"github.com/dilanEspindola/restapiFiber/database"
	"github.com/dilanEspindola/restapiFiber/routes"
	"github.com/gofiber/fiber/v2"
	// "math/rand"
	// "strconv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/users", routes.GetUsers)
	app.Get("api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("api/users/:id", routes.DeleteUser)
}

func main() {
	database.DbConnection()
	app := fiber.New()

	setUpRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hola")
	})

	app.Listen(":4000")
}
