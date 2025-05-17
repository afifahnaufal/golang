package routes

import (
	"github.com/gofiber/fiber/v2"
)

// IndexRoute mengatur semua routing aplikasi
func IndexRoute(app *fiber.App) {
	// Route default HTML render
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Homepage",
			"Name":  "Afifah Naufal",
		})
	})

	// Route about
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("📄 Terima kasih sudah mengunjungi halaman ini.")
	})

	// Route dengan parameter URL
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("👋 Hello " + name + "! Semoga harimu menyenangkan!")
	})

	// Route JSON API sederhana
	app.Get("/api/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"nama":    "Afifah Naufal",
			"jurusan": "D4 Teknik Informatika",
			"hobi":    []string{"makan dimsum", "menyanyi"},
		})
	})

	// Route POST sederhana (Body JSON)
	app.Post("/submit", func(c *fiber.Ctx) error {
		type Request struct {
			Nama string `json:"nama"`
		}

		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Ups Data tidak valid")
		}

		return c.SendString("Hello, " + req.Nama + "! Terima kasih sudah submit!")
	})

	// Route dengan query parameter
	app.Get("/cari", func(c *fiber.Ctx) error {
		keyword := c.Query("q")
		if keyword == "" {
			return c.SendString("Silakan masukkan parameter ?q=sesuatu")
		}
		return c.SendString("Kamu mencari: " + keyword)
	})
}
