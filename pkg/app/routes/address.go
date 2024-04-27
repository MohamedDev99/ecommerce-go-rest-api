package routes

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/gofiber/fiber/v3"
)

type AddressRouter struct {
	addressController controllers.AddressController
}

// initialize address routes
func NewAddressRouter(addressController controllers.AddressController) *AddressRouter {
	// initialize address router
	return &AddressRouter{
		addressController: addressController,
	}
}

// setup address routes
func (a *AddressRouter) Setup(addressRouterGroup fiber.Router) {
	// routes
	addressRoutes := addressRouterGroup.Group("/address")
	// 1- create address
	addressRoutes.Post("/", a.addressController.CreateAddress)
	// 2- delete address
	addressRoutes.Delete("/", a.addressController.DeleteAddress)
	// 3- edit home address
	addressRoutes.Put("/home", a.addressController.EditHomeAddress)
	// 4- edit work address
	addressRoutes.Put("/work", a.addressController.EditWorkAddress)
	log.Info("-----> ✅ Finish address routes ... ")
}
