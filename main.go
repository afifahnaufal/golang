package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/routes" // Ganti "gofiber" sesuai nama module di go.mod kamu
)

func main() {
	app := fiber.New()

	app.Get("/", routes.Home)
	app.Get("/about", routes.About)

	app.Listen(":3000")
}
