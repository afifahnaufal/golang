package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gofiber/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New()) // middleware logger

	routes.IndexRoute(app)

	app.Listen(":3000")
}
