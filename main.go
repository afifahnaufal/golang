package main

import (
	"gofiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Inisialisasi template engine
	engine := html.New("./views", ".html")

	// Buat instance Fiber dan hubungkan dengan engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Panggil semua routing dari routes
	routes.IndexRoute(app)

	// Jalankan server
	app.Listen(":3000")
}
