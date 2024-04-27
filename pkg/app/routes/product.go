package routes

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/controllers"
	"github.com/gofiber/fiber/v3"
)

type ProductRouter struct {
	ProductController controllers.ProductController
}

// new product router
func NewProductRouter(productController controllers.ProductController) *ProductRouter {
	// initialize product router
	return &ProductRouter{
		ProductController: productController,
	}
}

// setup product routes
func (p *ProductRouter) Setup(productRouterGroup fiber.Router) {
	// routes
	productRoutes := productRouterGroup.Group("/products")
	// 1- product viewer by admin
	productRoutes.Post("/admin", p.ProductController.ProductViewerByAdmin)
	// 2- search product
	productRoutes.Get("/search", p.ProductController.SearchProduct)
	// 3- search product by query
	productRoutes.Get("/search/query", p.ProductController.SearchProductByQuery)
	log.Info("-----> ✅ Finish product routes ... ")
}
