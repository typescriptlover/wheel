package schema

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var validate = validator.New()

func ValidateSchema(s interface{}) []*ErrorResponse {
	var errs []*ErrorResponse

	if err := validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var res ErrorResponse

			res.Field = err.Field()
			res.Tag = err.Tag()
			res.Value = err.Param()

			errs = append(errs, &res)
		}
	}

	return errs
}
