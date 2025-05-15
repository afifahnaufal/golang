package main

import (
    "github.com/gofiber/fiber/v2"
    "gofiber/routes" // gunakan nama module kamu (lihat di go.mod)
)

func main() {
    app := fiber.New()

    routes.SetupRoutes(app) // memuat semua route dari folder routes

    app.Listen(":3000")
}
