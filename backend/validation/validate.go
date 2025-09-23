package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("passwd", ValidatePassword)
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var re = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`)
	return re.MatchString(password)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
