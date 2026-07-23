package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecoveryJSON() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("panic recovered: %v", r)
					_ = c.JSON(http.StatusInternalServerError, echo.Map{
						"error": "internal_error",
					})
				}
			}()
			return next(c)
		}
	}
}
