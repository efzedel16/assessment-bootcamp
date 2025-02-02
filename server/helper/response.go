package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	apiResponse := Response{
		Meta: meta,
		Data: data,
	}

	return apiResponse
}

func ErrorResponse(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}

	return errors
}
