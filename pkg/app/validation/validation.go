package validation

import (
	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	types "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs/types/in"
	"github.com/go-playground/validator/v10"
)

var (
	log = configs.Logger
	// initialize validation
	validate = validator.New()
)

// validate : to validate the body, parameters and query
func Validate(data interface{}) []types.ValidationError {
	validationErrors := []types.ValidationError{}

	// validate data
	err := validate.Struct(data)

	if err != nil {
		// handle validation errors
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, types.ValidationError{
				Error:       true,
				ErrorMsg:    err.Error(),
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}

		return validationErrors
	}

	return validationErrors
}
