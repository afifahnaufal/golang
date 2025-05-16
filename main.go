package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gofiber/routes" // Pastikan nama modul di go.mod adalah "gofiber"
)

func main() {
	// Inisialisasi template engine (folder views, ekstensi .html)
	engine := html.New("./views", ".html")

	// Buat app Fiber dengan konfigurasi engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routing
	routes.IndexRoute(app)

	// Jalankan server di port 3000
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
