package controllers

import (
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/out"
	"github.com/gofiber/fiber/v3"
)

type CartController struct {
	cartRepo types.CartRepository
}

func NewCartController(cartRepo types.CartRepository) *CartController {
	// initialize cart controller
	return &CartController{cartRepo: cartRepo}
}

// add to cart
func (cart *CartController) AddToCart(c fiber.Ctx) error {
	// get product
	return c.SendStatus(200)
}

// remove item from cart
func (cart *CartController) RemoveItemFromCart(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// get item from cart
func (cart *CartController) GetItemFromCart(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// buy from cart
func (cart *CartController) BuyFromCart(c fiber.Ctx) error {
	return c.SendStatus(200)
}

// instant buy
func (cart *CartController) InstantBuy(c fiber.Ctx) error {
	return c.SendStatus(200)
}
