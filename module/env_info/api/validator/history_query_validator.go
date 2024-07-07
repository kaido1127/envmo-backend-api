package validator_module

import (
	"strings"

	"envmo/common/errors/error_response"

	"github.com/go-playground/validator/v10"
)

type HistoryQueryValidator struct {
	Valid *validator.Validate
}

func NewHistoryQueryValidator() *HistoryQueryValidator {
	v := &HistoryQueryValidator{Valid: validator.New()}
	return v
}

type HistoryQuery struct {
	DeviceID  string `query:"device_id" validate:"required"`
	StartTime int64  `query:"start_time" validate:"required,gte=1"`
	EndTime   int64  `query:"end_time" validate:"required,gte=1"`
}

func (v *HistoryQueryValidator) Validate(data interface{}) *error_response.ErrorResponse {
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
