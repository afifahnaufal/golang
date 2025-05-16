package main

import (
	"gofiber/routes" // ini harus sesuai module name di go.mod

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// panggil semua route dari file index.go
	routes.IndexRoute(app)

	app.Listen(":3000")
}
