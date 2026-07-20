package main

import (
	"log"
	"net/http"
	"task_tracker/internal/model"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)


func main() {
    e := echo.New()

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Task{})

	e.GET("/health", func(c echo.Context) error {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": "error",
		})
	}
		return c.JSON(http.StatusOK, echo.Map{
			"status": "everything is OK",
	})
	})
	e.Logger.Fatal(e.Start(":8080"))
}