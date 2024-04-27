package responses

import "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"

// address success response with address data
type AddressResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Address models.Address `json:"address"`
}

// address success response
type AddressSuccess struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// address fail response
type AddressFail struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// address error response
type AddressErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	AddressSuccessStatus = "success"
	AddressFailStatus    = "fail"
	AddressErrorStatus   = "error"
	// error messages
	ErrAddressNotFound       = "address not found"
	ErrWhileInsertingAddress = "error while inserting address"
	ErrWhileUpdatingAddress  = "error while updating address"
	ErrWhileDeletingAddress  = "error while deleting address"
	ErrWhileGettingAddress   = "error while getting address"
	ErrWhileGettingAddresses = "error while getting addresses"
)
