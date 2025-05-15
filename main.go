package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware logger dari Fiber
	app.Use(logger.New())

	// Route utama (home)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("👋 Hello, World! Selamat datang di halaman ini!! Ini halaman tutorial 😊")
	})

	// Route about (cukup satu)
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("ℹ️ Ini halaman about! Terima kasih sudah mengunjungi.")
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

	// Route contact
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("📞 Hubungi kami di: 123-456-7890 atau email: contoh@email.com")
	})

	app.Listen(":3000")
}
