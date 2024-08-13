package routes

import (
	usercontrollers "github.com/fadhilmufid/fiber-tutorial/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})
	
	// User
	app.Get("/users/:id", usercontrollers.Show)
}