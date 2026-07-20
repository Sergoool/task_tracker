package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

    type Task struct {
		ID 				string `json:"id"`
		Title			string `json:"title"` 
		Description		string `json:"description"` 
		Status			string `json:"status"` 
		CreatedAt		string `json:"created_at"`
		UpdatedAt		string `json:"updated_at"`
	}


func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Добро пожаловать")
    })

    // Запускаем сервер на порту 8080
    e.Start(":8080")
}