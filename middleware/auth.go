package middleware

import (
	"net/http"

	"github.com/rorikurniadi/simple-auth/models"

	"github.com/asaskevich/govalidator"
)

// Validation Input
func AuthValidate(user models.User, res http.ResponseWriter, req *http.Request) (string, bool) {
	if _, errValid := govalidator.ValidateStruct(user); errValid != nil {
		return errValid.Error(), false
	}

	email := govalidator.IsEmail(user.Email)
	if !email {
		return "Email is required", false
	}

	if user.Password == "" {
		return "Password is required", false
	}

	return "validation success", true
}
