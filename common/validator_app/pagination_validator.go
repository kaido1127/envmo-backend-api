package validator_app

import (
	"strings"

	"envmo/common/errors/error_response"

	"github.com/go-playground/validator/v10"
)

type PaginationValidator struct {
	Valid *validator.Validate
}

func NewPaginationValidator() *PaginationValidator {
	v := &PaginationValidator{Valid: validator.New()}
	return v
}

type PaginationQuery struct {
    Limit    int    `query:"limit" validate:"required,gte=1,lte=100"`
	Page     int    `query:"page" validate:"required,gte=1"`
}

func (v *PaginationValidator) Validate(data interface{}) *error_response.ErrorResponse {
	if err := v.Valid.Struct(data); err != nil {
		var errMsg strings.Builder
		errMsg.WriteString("Invalid input:")
		for _, err := range err.(validator.ValidationErrors) {
			errMsg.WriteString(" " + err.Field() + ": " + err.Tag())
		}
		return error_response.BadRequestErrorResponse(errMsg.String())
	}
	return nil
}


