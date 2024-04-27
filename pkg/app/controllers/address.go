package controllers

import (
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/out"
	"github.com/gofiber/fiber/v3"
)

type AddressController struct {
	addressRepository types.AddressRepository
}

func NewAddressController(addressRepository types.AddressRepository) *AddressController {
	// create new address controller
	return &AddressController{addressRepository: addressRepository}
}

// create address
func (ac *AddressController) CreateAddress(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// edit home address
func (ac *AddressController) EditHomeAddress(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// edit work address
func (ac *AddressController) EditWorkAddress(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// delete address
func (ac *AddressController) DeleteAddress(c fiber.Ctx) error {
	return c.SendStatus(200)
}
