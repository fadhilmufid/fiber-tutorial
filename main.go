package main

import (
	"log"

	"github.com/fadhilmufid/fiber-tutorial/database"
	"github.com/fadhilmufid/fiber-tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)


func main() {
	database.ConnectDb()
	// Engine
	engine := html.New("./app/views", ".html")
	// App
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// static
	app.Static("/", "./app/assets")


	routes.Register(app)

	log.Fatal(app.Listen(":3000"))
}
