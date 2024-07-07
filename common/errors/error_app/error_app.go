package error_app

import "errors"

const (
	duplicateKeyErrorCode     = "duplicate-key-error"
	documentNotFoundErrorCode = "document-not-found-error"
	invalidJwtTokenErrorCode  = "invalid-jwt-token-error"
	expiredJwtTokenErrorCode  = "expried-jwt-token-error"
	unauthorizedErrorCode     = "unauthorized-error"
	permissionDeniedErrorCode = "permission-denied-error"
	documentFormatErrorCode   = "document-format-error"

)

var ErrMongoDocumentNotFound = errors.New(documentNotFoundErrorCode)
var ErrInvalidJwtToken = errors.New(invalidJwtTokenErrorCode)
var ErrExpriedJwtToken = errors.New(expiredJwtTokenErrorCode)
var ErrDuplicateKey = errors.New(duplicateKeyErrorCode)
var ErrUnauthorized = errors.New(unauthorizedErrorCode)
var ErrDocumentFormat = errors.New(documentFormatErrorCode)
var ErrPermissionDenied = errors.New(permissionDeniedErrorCode)
