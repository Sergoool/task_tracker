package main

import (
	"log"
	"net/http"
	"task_tracker/internal/api/rest/handlers"
	"task_tracker/internal/config"
	"task_tracker/internal/connection/initialize"
	"task_tracker/internal/domain/repository"
	"task_tracker/internal/domain/service"
	"task_tracker/internal/domain/types"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"os"
)


func main() {
	_ = godotenv.Load()

	cfg, err := config.Load() 
	if err != nil {         
		log.Fatalf("config error: %v", err) 
	}

	gormDB, err := initialize.New(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("database error: %v", err)
	}
	log.Println("database connected")

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("db sql error: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("db ping error: %v", err)
	}
	log.Println("db ping ok") 

	if err := gormDB.AutoMigrate(&types.Task{}); err != nil {
		log.Fatalf("db migrate error: %v", err)
	}
	log.Println("db migrated")

    e := echo.New()

	taskRepo := repository.NewTaskGormRepository(gormDB)
	taskService := service.NewTaskService(taskRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	versionHandler := handlers.NewVersionHandler(taskService)

	api := e.Group("/api")
	{
		api.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map {
				"status": "ok",
			})
		})
		api.GET("/ping", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map {
				"message": "pong",
			})
		})
		api.GET("/version", versionHandler.GetVersion)
		api.POST("/tasks", taskHandler.Create)
		api.GET("/tasks", taskHandler.List)
		api.GET("/tasks/:id", taskHandler.GetByID)
		api.PATCH("/tasks/:id", taskHandler.Update)
		api.DELETE("/tasks/:id", taskHandler.Delete)
	}

	_ = gormDB

	addr := ":" + cfg.Port      
	log.Printf("starting server on %s", addr) 

	if err := e.Start(addr); err != nil { 
		log.Fatalf("server error: %v", err)
	}
}