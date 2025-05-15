package routes

import (
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Selamat datang di halaman utama 😄")
    })

    app.Get("/about", func(c *fiber.Ctx) error {
        return c.SendString("Ini halaman tentang (about) 😎")
    })
}
