package handlers

import (
	"github.com/gofiber/fiber/v2"
	Access "touristapp.com/controllers/access"
)

//GetAllAccesses fetches all the accesses available
func GetAllAccesses(c *fiber.Ctx) error {

	accesses, err := Access.GetAll()
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 200,
		"data": accesses,
	})
}
