package types

import "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"

type ProductRepository interface {
	// CreateProduct(product models.Product) error
	// GetProduct(id string) (models.Product, error)
	ProductViewerByAdmin(id string) (models.Product, error)
	SearchProduct() ([]models.Product, error)
	SearchProductByQuery(query string) ([]models.Product, error)
	// GetProducts() ([]models.Product, error)
	// UpdateProduct(product models.Product) error
	// DeleteProduct(id string) error
}
