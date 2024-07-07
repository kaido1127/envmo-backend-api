package error_response

import "envmo/common/errors/error_app"

type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

const (
	notFoundErrCode       = "not-found-error"
	unauthorizedErrCode   = "unauthorized-error"
	internalServerErrCode = "internal-server-error"
	badRequestErrCode     = "bad-request-error"
	conflictErrCode       = "conflict-error"
)

func NewErrorResponse(errorCode string, message string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func BadRequestErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(badRequestErrCode, message)
}

func NotFoundErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(notFoundErrCode, message)
}

func UnauthorizedErroResponse(message string) *ErrorResponse {
	return NewErrorResponse(unauthorizedErrCode, message)
}

func InternalServerErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(internalServerErrCode, message)
}

func ConflictErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(conflictErrCode, message)
}

func PermissionDeniedErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(error_app.ErrPermissionDenied.Error(), message)
}

func InvalidJwtTokenErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(error_app.ErrInvalidJwtToken.Error(), message)
}

func ExpriedJwtTokenErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(error_app.ErrExpriedJwtToken.Error(), message)
}
