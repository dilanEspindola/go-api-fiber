package main

import (
	"github.com/dilanEspindola/restapiFiber/routes"
	"github.com/gofiber/fiber/v2"
	// "math/rand"
	// "strconv"
)

func main() {
	app := fiber.New()

	app.Get("/", routes.GetUsers)
	app.Get("/:id", routes.GetUser)

	app.Listen(":4000")
}
