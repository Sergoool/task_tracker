package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func ErrorLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
				err :=next(c)
		if err != nil {
			log.Printf("request error method=%s path=%s status=%d err=%v",
				c.Request().Method,
				c.Request().URL.Path,
				c.Response().Status,
				err,
			)
			return err
		}
return nil
		}
	}
}