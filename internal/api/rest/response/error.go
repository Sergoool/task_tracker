package response

import (
	"errors"
	"net/http"

	"task_tracker/internal/domain/service"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
}

func JSONError(c echo.Context, status int, code string, details any) error {

	return c.JSON(status, ErrorResponse{
		Error:   code,
		Details: details,
	})
}

func FromServiceError(c echo.Context, err error) error {
	var appErr *service.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case service.CodeValidation:
			return JSONError(c, http.StatusBadRequest, string(appErr.Code), appErr.Details) // 400
		case service.CodeNotFound:
			return JSONError(c, http.StatusNotFound, string(appErr.Code), appErr.Details) // 404
			
		default:
			return JSONError(c, http.StatusInternalServerError, string(service.CodeInternal), nil) // 500
			
		}
	}


	return JSONError(c, http.StatusInternalServerError, string(service.CodeInternal), nil) // 500
}
