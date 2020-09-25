package routers

import (
	"github.com/gofiber/fiber/v2"
	"touristapp.com/handlers"
)

//Routes defines all the routings in this api
func Routes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	api.Get("/accesses", handlers.GetAllAccesses)

	bga := api.Group("/BGA")
	aga := api.Group("/AGA")

	aga.Get("/services", handlers.GetAllServices)
	aga.Post("/services", handlers.AddService)
	aga.Post("/services/price", handlers.AddPriceDetails)
	aga.Post("/services/supplier", handlers.MapSupplier)
	aga.Post("/services/tc", handlers.MapTC)
	aga.Get("/suppliers", handlers.GetAllSuppliers)
	aga.Post("/suppliers", handlers.AddSupplier)
	aga.Get("/placeholder_services", handlers.GetAllPlaceholderServices)
	aga.Get("/99a_services", handlers.GetAll99AServices)

	bga.Get("/services", handlers.GetAllServices)
	bga.Get("/agents", handlers.GetAllAgents)
	bga.Get("/agents/:id", handlers.GetAgent)
	bga.Post("/agents", handlers.AddAgent)
	bga.Get("/sub_agents", handlers.GetAllSubAgents)
	bga.Get("/sub_agents/:id", handlers.GetSubAgent)
}
