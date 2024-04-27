package responses

import "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"

// product success response
type ProductSuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// product success response with product data
type ProductSuccessResponseWithProductData struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Product models.Product `json:"product"`
}

// product fail response
type ProductFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// product error response
type ProductErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	ProductCreateSuccess = "product created successfully"
	ProductUpdateSuccess = "product updated successfully"
	ProductDeleteSuccess = "product deleted successfully"
	// fail errors
	ErrProductNotFound       = "product not found"
	ErrProductAlreadyExists  = "product already exists"
	ErrProductNotCreated     = "product not created"
	ErrProductNotUpdated     = "product not updated"
	ErrProductNotDeleted     = "product not deleted"
	ErrProductIdIsNotValid   = "product id is not valid"
	ErrProductNotValid       = "product not valid"
	ErrProductNotFoundInCart = "product not found in cart"
	ErrWhileGettingProducts  = "error while getting products"
	ErrWhileGettingProduct   = "error while getting product"
	ErrWhileInsertingProduct = "error while inserting product"
	ErrWhileUpdatingProduct  = "error while updating product"
	ErrWhileDeletingProduct  = "error while deleting product"
)
