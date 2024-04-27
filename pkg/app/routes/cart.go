package routes

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/gofiber/fiber/v3"
)

type CartRouter struct {
	cartController controllers.CartController
}

// initialize cart routes
func NewCartRouter(cartController controllers.CartController) *CartRouter {
	// initialize cart router
	return &CartRouter{
		cartController: cartController,
	}
}

// setup cart routes
func (c *CartRouter) Setup(cartRouterGroup fiber.Router) {
	// routes
	cartRoutes := cartRouterGroup.Group("/carts")
	// 1- add to cart
	cartRoutes.Post("/", c.cartController.AddToCart)
	// 2- remove item from cart
	cartRoutes.Delete("/", c.cartController.RemoveItemFromCart)
	// 3- get item from cart
	cartRoutes.Get("/", c.cartController.GetItemFromCart)
	// 4- buy from cart
	cartRoutes.Post("/buy", c.cartController.BuyFromCart)
	// 5- instant buy
	cartRoutes.Post("/buy/instant", c.cartController.InstantBuy)
	log.Info("-----> ✅ Finish cart routes ... ")
}
