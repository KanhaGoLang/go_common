package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
)

func HandleError(w http.ResponseWriter, err error, status int) {
	MyLogger.Println(color.RedString(err.Error()))
	http.Error(w, err.Error(), status)
}

func HandleValidationErrors(w http.ResponseWriter, err error) {
	var validationErrors = make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors[GetJSONTag(err)] = GetValidationErrorMsg(err)
	}

	errorResponse := ErrorResponse{Message: "Validation failed", Details: validationErrors}

	MyLogger.Println(color.RedString("ErrorResponse %v", errorResponse))

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}

func GetJSONTag(err validator.FieldError) string {
	field, _ := reflect.TypeOf(User{}).FieldByName(err.StructField())
	return field.Tag.Get("json")
}

func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return len(password) >= 8
}

func ValidateRole(roles []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		for _, validRole := range roles {
			if validRole == role {
				return true
			}
		}

		return false
	}

}

func GetValidationErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s is below minimum length", err.Field())
	case "max":
		return fmt.Sprintf("%s is above maximum length", err.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email address", err.Field())
	case "validateRole":
		return fmt.Sprintf("%s is not one of the allowed values %s", err.Field(), strings.Join(ValidRoles, ", "))
	case "strength":
		return fmt.Sprintf("%s is too weak", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
