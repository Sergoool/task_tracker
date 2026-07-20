package main

import (
	"net/http"
	"task_tracker/internal/model"

	"github.com/labstack/echo/v4"
)


func main() {
    e := echo.New()
	e.GET("/", func(c echo.Context) error {
		task := model.Task {
			ID: "1",
			Title: "Example",
			Description: "work",
			Status: model.New,
			CreatedAt: "2026-07-20",
			UpdatedAt: "2026-07-20",
		}
		return c.JSON(http.StatusOK, task)
	})
	e.Logger.Fatal(e.Start(":8080"))
}