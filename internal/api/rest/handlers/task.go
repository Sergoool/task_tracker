package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"task_tracker/internal/api/rest/dto"
	"task_tracker/internal/domain/service"
	"task_tracker/internal/domain/types"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func toTaskResponse(t any) dto.TaskResponse {
	return dto.TaskResponse{} 
}

func (h *TaskHandler) Create(c echo.Context) error {
	var req dto.CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
		
	}

	task, err := h.taskService.Create(c.Request().Context(), req.Title, req.Description)
	if err != nil {
		if errors.Is(err, types.ErrValidation) {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "validation error"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal error"})
	}

	return c.JSON(http.StatusCreated, dto.TaskResponse{
		ID:        task.ID, 
		Title:     task.Title,
		Description: task.Description,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	})
}

func (h *TaskHandler) List(c echo.Context) error {
	tasks, err := h.taskService.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal error"})
	}

	resp := make([]dto.TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, dto.TaskResponse{
			ID:        t.ID,
			Title:     t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt: t.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) GetByID(c echo.Context) error { // GET /tasks/id
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id64 == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	task, err := h.taskService.GetByID(c.Request().Context(), uint(id64))
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal error"})
	}

	return c.JSON(http.StatusOK, dto.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Description: task.Description,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	})
}


