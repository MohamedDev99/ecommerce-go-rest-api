package responses

import (
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/in"
)

// success response
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// fail or error response
type FailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// error response

// validation response
type ValidationResponse struct {
	Status           string                  `json:"status"`
	TypeOfValidation string                  `json:"typeOfValidation"`
	ValidationErrors []types.ValidationError `json:"validationErrors"`
}

const (
	// status
	SuccessStatus = "success"
	FailStatus    = "fail"
	ErrorStatus   = "error"
	// error messages
	ErrInternalServerError = "internal server error"
)
