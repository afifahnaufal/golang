package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Selamat datang di halaman ini!! Ini halaman tutorial")
	})

	// Tambahan untuk commit ke-2
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("Ini halaman about!")
	})

	app.Listen(":3000")
}
