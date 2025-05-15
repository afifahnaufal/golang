package routes

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("🏡 Ini halaman utama! Selamat datang ya!")
}

func About(c *fiber.Ctx) error {
	return c.SendString("📖 Ini halaman About. Dibuat oleh Afifah.")
}
