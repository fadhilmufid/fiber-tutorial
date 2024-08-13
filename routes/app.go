package routes

import (
	"github.com/fadhilmufid/fiber-tutorial/app/features/user"
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
	app.Get("/users/:id", user.ShowController)
	app.Post("/users", user.CreateController)
	app.Get("/users", user.IndexController)
}