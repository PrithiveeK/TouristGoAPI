package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	Service "touristapp.com/controllers/service"
	Models "touristapp.com/models"
)

//GetAllServices fetches all the services in the database
func GetAllServices(c *fiber.Ctx) error {

	services, err := Service.GetAll(map[string]string{
		"type": "",
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 200,
		"data": services,
	})
}

//AddService inserts a new service into the database
func AddService(c *fiber.Ctx) error {
	var newService struct {
		Service        Models.NewService `json:"services" form:"services"`
		Categories     []int64           `json:"categories" form:"categories"`
		Amenities      []int64           `json:"amenities" form:"amenities"`
		LinkedServices []int64           `json:"linkedServices" form:"linkedServices"`
	}

	if err := c.BodyParser(&newService); err != nil {
		log.Printf("Invalid service data: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}

	service, err := Service.Add(newService.Service, map[string]string{
		"type": "",
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong",
		})
	}

	if err := Service.MapCategories(service.ID, newService.Categories); err != nil {
		log.Printf("Error mappinng categories to services")
	}

	if err := Service.MapAmenities(service.ID, newService.Amenities); err != nil {
		log.Printf("Error mappinng categories to services")
	}

	if err := Service.MapLinkedServices(service.ID, newService.LinkedServices); err != nil {
		log.Printf("Error mappinng categories to services")
	}

	return c.JSON(&fiber.Map{
		"code": 201,
		"data": service,
		"msg":  "Successfully Added!",
	})
}

//AddPriceDetails adds the price details for the service
func AddPriceDetails(c *fiber.Ctx) error {
	var newPD struct {
		ServiceID int64                    `json:"service_id" form:"service_id"`
		Price     []Models.NewPriceDetails `json:"price" form:"price"`
	}
	if err := c.BodyParser(&newPD); err != nil {
		log.Printf("Invalid PriceDetails: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}
	if err := Service.AddPricing(newPD.ServiceID, newPD.Price); err != nil {
		log.Printf("Error adding price details to services")
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 201,
		"data": nil,
		"msg":  "Successfully Inserted",
	})
}

//MapSupplier maps the supplier to the service
func MapSupplier(c *fiber.Ctx) error {
	var newSupplier struct {
		ServiceID  int64 `json:"service_id" form:"service_id"`
		SupplierID int64 `json:"supplier_id" form:"supplier_id"`
	}
	if err := c.BodyParser(&newSupplier); err != nil {
		log.Printf("Invalid Supplier: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}
	if err := Service.MapSupplier(newSupplier.ServiceID, newSupplier.SupplierID); err != nil {
		log.Printf("Error mappinng supplier to services")
	}
	return c.JSON(&fiber.Map{
		"code": 201,
		"data": nil,
		"msg":  "Successfully Mapped",
	})
}

//MapTC maps the supplier to the service
func MapTC(c *fiber.Ctx) error {
	var newTC struct {
		ServiceID int64 `json:"service_id" form:"service_id"`
		TCID      int64 `json:"tc_id" form:"tc_id"`
	}
	if err := c.BodyParser(&newTC); err != nil {
		log.Printf("Invalid tc: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}
	if err := Service.MapTC(newTC.ServiceID, newTC.TCID); err != nil {
		log.Printf("Error mappinng tc to services")
	}
	return c.JSON(&fiber.Map{
		"code": 201,
		"data": nil,
		"msg":  "Successfully Mapped",
	})
}

//GetAll99AServices fetches all the services which are 99a
func GetAll99AServices(c *fiber.Ctx) error {

	services, err := Service.GetAll(map[string]string{
		"type": "99A",
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 200,
		"data": services,
	})
}

//GetAllPlaceholderServices fetches all the services which are placeholders
func GetAllPlaceholderServices(c *fiber.Ctx) error {

	services, err := Service.GetAll(map[string]string{
		"type": "placeholder",
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 200,
		"data": services,
	})
}
