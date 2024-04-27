package validation

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/app/responses"
	"github.com/gofiber/fiber/v3"
)

type RequestAuthLoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// validate login body
func ValidateLoginBody(c fiber.Ctx) (*RequestAuthLoginBody, interface{}) {
	// body
	loginBody := new(RequestAuthLoginBody)

	if err := c.Bind().JSON(loginBody); err != nil {
		log.Info("validation of login body failed")
		return &RequestAuthLoginBody{}, responses.FailResponse{Status: "fail", Message: err.Error()}
	}

	// validation of body
	// check if body is have email and password
	if validationErrors := Validate(loginBody); len(validationErrors) > 0 {
		log.Info("validation of login body failed")
		return &RequestAuthLoginBody{}, responses.ValidationResponse{Status: "fail", TypeOfValidation: "login", ValidationErrors: validationErrors}
	}

	return loginBody, nil
}

type RequestAuthRegisterBody struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,alphanum,min=2,max=20"`
	LastName  string `json:"lastName" validate:"required,alphanum,min=2,max=20"`
	Password  string `json:"password" validate:"required,min=8"`
	Phone     string `json:"phone" validate:"required,numeric,min=10,max=10"`
	Role      string `json:"role" validate:"required,oneof=admin user"`
}

// validate register body
func ValidateRegisterBody(c fiber.Ctx) (*RequestAuthRegisterBody, interface{}) {
	// body
	registerBody := new(RequestAuthRegisterBody)

	log.Info("validation of register body 1")

	if err := c.Bind().JSON(registerBody); err != nil {
		log.Info("validation of register body failed")
		return &RequestAuthRegisterBody{}, responses.FailResponse{Status: "fail", Message: err.Error()}
	}
	log.Info("validation of register body 2")

	// validation of body
	if validationErrors := Validate(registerBody); len(validationErrors) > 0 {
		log.Info("validation of register body failed")
		return &RequestAuthRegisterBody{}, responses.ValidationResponse{Status: "fail", TypeOfValidation: "register", ValidationErrors: validationErrors}
	}

	return registerBody, nil
}
