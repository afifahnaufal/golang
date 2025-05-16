package routes

import "github.com/gofiber/fiber/v2"

func IndexRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Selamat datang di halaman utama!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("📄 Terima kasih sudah mengunjungi halaman ini.")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("👋 Hallo " + name + "! Semoga harimu menyenangkan!")
	})
}
