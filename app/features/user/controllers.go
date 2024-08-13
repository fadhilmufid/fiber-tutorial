package user

import (
	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	users := FindAll()
	responseUsers := SerializeAll(users)
	return c.Status(200).JSON(responseUsers)
}


func ShowController(c *fiber.Ctx) error {
	// Validate Integer
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	// Validate User
	user, err := FindById(id)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := Serialize(user)
	return c.Render(
				"index",
	 			fiber.Map{
					"Title": responseUser,
				},
			 	"layouts/main")

}


func CreateController(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	CreateData(&user)

	responseUser := Serialize(user)
	return c.Status(200).JSON(responseUser)
}


