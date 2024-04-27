package responses

// cart success response
type CartSuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// cart fail response
type CartFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// cart error response
type CartErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	CartCreateSuccess = "Cart created successfully"
	CartUpdateSuccess = "Cart updated successfully"
	CartDeleteSuccess = "Cart deleted successfully"
	ErrCartNotFound   = "Cart not found"
	ErrCartExists     = "Cart already exists"
	// error can't decode products
	ErrCantDecodeProducts = "Cant't decode products"
	// error can't remove item from cart
	ErrCantRemoveItemFromCart = "Cant't remove item from cart"
	// error can't get item from cart
	ErrCantGetItemFromCart = "Cant't get item from cart"
	// error can't buy from cart item
	ErrCantBuyCartItem    = "Cant buy cart item"
	ErrWhileDeletingCart  = "Error while deleting cart"
	ErrWhileGettingCart   = "Error while getting cart"
	ErrWhileInsertingCart = "Error while inserting cart"
)
