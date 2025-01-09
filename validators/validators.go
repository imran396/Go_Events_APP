package validators

import (
	"fmt"
	"web/model"
	"github.com/go-playground/validator/v10"
)

func  ValidateUser(user model.User) (map[string]interface{}, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	var jsonResponse map[string]interface{}
	if err != nil {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
	}
	errorMessages := []map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, map[string]string{
			"field":   err.Field(),
			"tag":     err.Tag(),
			"value":   fmt.Sprintf("%v", err.Value()),
			"message": fmt.Sprintf("Field '%s' failed on the '%s' validation", err.Field(), err.Tag()),
		})
	}
	jsonResponse = map[string]interface{}{
		"success": false,
		"errors":  errorMessages,
	}
  }
  return jsonResponse, err
}