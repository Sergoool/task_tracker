package handlers

import (
	"net/http"

	"task_tracker/internal/domain/service"

	"github.com/labstack/echo/v4"
)

type VersionHandler struct {
	taskService *service.TaskService
}

func NewVersionHandler(taskService *service.TaskService) *VersionHandler {
	return &VersionHandler{taskService: taskService}
}

func (h *VersionHandler) GetVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"version": h.taskService.Version(),
	})
}
