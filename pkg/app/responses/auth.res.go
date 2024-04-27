package responses

import "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"

// auth success response
type AuthSuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// auth success response with user
type AuthSuccessResponseWithUser struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

// auth fail response
type AuthFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// auth error response
type AuthErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	// user logged in successfully
	AuthSuccess = "user logged in successfully"
	// user registered successfully
	AuthRegisterSuccess = "user registered successfully"
	// user logged out successfully
	AuthLogoutSuccess = "user logged out successfully"
	// invalid email or password
	ErrAuthInvalidEmailOrPassword = "invalid email or password"
)
