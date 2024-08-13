package main

import (
	"log"

	"github.com/fadhilmufid/fiber-tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Engine
	engine := html.New("./app/views", ".html")

	// App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Register(app)

	app.Static("/", "./app/assets")
	log.Fatal(app.Listen(":3000"))
}
