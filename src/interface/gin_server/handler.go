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

func (g GinServer) sendError(ctx gin.Context, status int, err error, errCode int) {
	ctx.JSON(status, ErrorResponse{
		Error:     err.Error(),
		ErrorCode: errCode,
	})
}

func (g GinServer) errorHandle(ctx gin.Context, err error) {
	unwrapErr := errors.Unwrap(err)
	if unwrapErr == nil {
		unwrapErr = err
	}

	switch unwrapErr {
	case use_case.ErrorItemNotFound:
		g.sendError(ctx, http.StatusNotFound, err, 1)
		break
	case use_case.ErrorFirstNameIsRequire:
		g.sendError(ctx, http.StatusBadRequest, err, 2)
		break
	case use_case.ErrorLastNameIsRequire:
		g.sendError(ctx, http.StatusBadRequest, err, 3)
		break
	case use_case.ErrorPhoneIsRequire:
		g.sendError(ctx, http.StatusBadRequest, err, 4)
		break
	case use_case.ErrorEmailIsRequire:
		g.sendError(ctx, http.StatusBadRequest, err, 5)
		break
	case use_case.ErrorAddressIsRequire:
		g.sendError(ctx, http.StatusBadRequest, err, 6)
		break
	case use_case.ErrorDuplicatePhone:
		g.sendError(ctx, http.StatusBadRequest, err, 7)
		break
	case use_case.ErrorDuplicateEmail:
		g.sendError(ctx, http.StatusBadRequest, err, 8)
		break
	case use_case.ErrorDataTransform:
		g.sendError(ctx, http.StatusBadRequest, err, 9)
		break
	default:
		g.sendError(ctx, http.StatusInternalServerError, err, 0)
		break
	}
}
