package user

import (
	"graphql_search/models"

	"github.com/go-playground/validator/v10"
)

func validateSignupRequest(user *models.UserDB, validate *validator.Validate) []validator.FieldError {
	var errs = make([]validator.FieldError, 0)

	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errs
		}

		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err)
		}
	}

	return errs
}
