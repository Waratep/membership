package gin_server

import (
	"errors"
	"net/http"

	"github.com/Waratep/membership/src/use_case"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

func (g GinServer) errorHandle(ctx gin.Context, err error) {
	unwrapErr := errors.Unwrap(err)
	if unwrapErr == nil {
		unwrapErr = err
	}

	switch unwrapErr {
	case use_case.ErrorItemNotFound:
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 1,
		})
		break
	case use_case.ErrorFirstNameIsRequire:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 2,
		})
		break
	case use_case.ErrorLastNameIsRequire:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 3,
		})
		break
	case use_case.ErrorPhoneIsRequire:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 4,
		})
		break
	case use_case.ErrorEmailIsRequire:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 5,
		})
		break
	case use_case.ErrorAddressIsRequire:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 6,
		})
		break
	case use_case.ErrorDuplicatePhone:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 7,
		})
		break
	case use_case.ErrorDuplicateEmail:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 8,
		})
		break
	case use_case.ErrorDataTransform:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 9,
		})
		break
	default:
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:     err.Error(),
			ErrorCode: 0,
		})
		break
	}
}
