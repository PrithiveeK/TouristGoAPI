package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	Supplier "touristapp.com/controllers/supplier"
	Models "touristapp.com/models"
)

//GetAllSuppliers fetches all the supplier companies available
func GetAllSuppliers(c *fiber.Ctx) error {

	suppliers, err := Supplier.GetAll()
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 200,
		"data": suppliers,
	})
}

//AddSupplier is the handler for inserting a new row of supplier
func AddSupplier(c *fiber.Ctx) error {

	var newSupplier Models.NewSupplier

	if err := c.BodyParser(&newSupplier); err != nil {
		log.Printf("Cannot parse supplier body: %s\n", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}
	supplier, err := Supplier.Add(newSupplier)
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Inspecting incomming data",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 201,
		"data": supplier,
	})
}
