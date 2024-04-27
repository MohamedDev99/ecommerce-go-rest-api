package controllers

import (
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/out"
	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	productRepo types.ProductRepository
}

// product controller
func NewProductController(productRepo types.ProductRepository) *ProductController {
	// initialize product controller
	return &ProductController{
		productRepo: productRepo,
	}
}

// product viewer by admin
func (p *ProductController) ProductViewerByAdmin(c fiber.Ctx) error {
	// get product
	return c.SendStatus(200)
}

// search product for user
func (p *ProductController) SearchProduct(c fiber.Ctx) error {
	// get product
	return c.SendStatus(200)
}

// search product by query (filter)
func (p *ProductController) SearchProductByQuery(c fiber.Ctx) error {
	// get product
	return c.SendStatus(200)
}
