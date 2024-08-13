package main

import (
	"log"

	"github.com/fadhilmufid/fiber-tutorial/database"
	"github.com/fadhilmufid/fiber-tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)


type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func setupRoutes(app *fiber.App) {

app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})
	// app.Get("/users/:name?", func(c *fiber.Ctx) error {
		// return c.Render("index", fiber.Map{
		// 	"Title": c.Params("name"),
		// }, "layouts/main")
	// })

	app.Get("/users/:id", routes.GetUser)
	app.Post("/users", routes.CreateUser)
}


func main() {
	database.ConnectDb()
	// Engine
	engine := html.New("./app/views", ".html")

	// App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)

	app.Static("/", "./app/assets")
	log.Fatal(app.Listen(":3000"))
}
