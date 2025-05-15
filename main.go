package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
)
// Route utama (home)
func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("👋 Hello, World! Selamat datang di halaman ini!! Ini halaman tutorial 😊")
	})

	// Route about
	app.Get("/about", func(c *fiber.Ctx) error {
    return c.SendString("ℹ️ Ini halaman about! Terima kasih sudah mengunjungi.")
	})
	
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("ℹ️ Terima kasih sudah mengunjungi halaman ini.")
	})
	// Route dengan parameter
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("👋 Halo " + name + "! Semoga harimu menyenangkan!")
	})
	// Route JSON response
	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Ini response JSON",
		})
	})

	app.Listen(":3000")
}
