package main

import (
	"github.com/gofiber/fiber/v2" 
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// Middleware logging dan recover
	app.Use(logger.New())
	app.Use(recover.New())

	// Route utama (home)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("👋 Hello, World! Selamat datang di halaman ini!! Ini halaman tutorial 😊")
	})

	// Route about (hindari duplikasi)
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("ℹ️ thankyou! for coming at this page.")
	})

	// Route dengan parameter
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("👋 Hallo " + name + "! Semoga harimu menyenangkan!")
	})

	// Route JSON response
	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "sudah di respon JSON",
		})
	})

	// Route contact
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("📞 Hubungi kami di: 123-456-7890 atau email: afifah@email.com")
	})

	app.Listen(":3000")
}
