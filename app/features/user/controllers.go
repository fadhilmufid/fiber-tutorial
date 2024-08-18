package user

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	users := FindAll() // Fetch the users from the database or service
    OldresponseUsers := SerializeAllJSON(users) // Get JSON string from users

    // Construct the new JSON response
    responseUsers := fmt.Sprintf(`{"status": 200, "users": %s}`, OldresponseUsers)

    // Define a variable to hold the unmarshaled data
    var responseData map[string]interface{}

    // Unmarshal the JSON string into the map
    err := json.Unmarshal([]byte(responseUsers), &responseData)
    if err != nil {
        log.Println("Error unmarshaling JSON:", err)
        return c.Status(500).SendString("Internal Server Error")
    }

    // Pass the unmarshaled data to the template
    return c.Render(
        "users/index",
        fiber.Map{
            "response": responseData, // Pass the unmarshaled data
        },
        "layouts/main")

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

	// responseUser := Serialize(user)
	
	responseUserJSON := SerializeJSON(user)

	return c.Render(
				"users/show",
	 			fiber.Map{
					"user": responseUserJSON,
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


