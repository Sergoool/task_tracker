package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"task_tracker/internal/api/rest/dto"
	"task_tracker/internal/api/rest/response"
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
		return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
			"json":"invalid",
		})
		
	}

	task, err := h.taskService.Create(c.Request().Context(), req.Title, req.Description)
	if err != nil {
			return response.FromServiceError(c, err)
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
	var statusPtr *string
	if statusStr := c.QueryParam("status"); statusStr != "" {
		statusPtr = &statusStr
	}
	limit := 20
	offset := 0

	if s := c.QueryParam("limit"); s != "" {
		v, err := strconv.Atoi(s)
		if err != nil {
			return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
				"error": "invalid limit",
			})
		}
		limit = v
	}
	if s := c.QueryParam("offset"); s != "" {
		v, err := strconv.Atoi(s)
		if err != nil || v < 0 {
			return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
				"error":"invalid offset",
			})
		}
		offset = v
	}
	tasks, err := h.taskService.List(c.Request().Context(), statusPtr, limit, offset)
	if err != nil {
		/*
		if errors.Is(err, types.ErrValidation) {
			return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
				"error": "validation error",
			})
		}
		return response.JSONError(c, http.StatusInternalServerError, "validation_error", map[string]string{
			"error": "internal error",
		})*/
		return response.FromServiceError(c, err)
	}


	resp := make([]dto.TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) GetByID(c echo.Context) error { // GET /tasks/id
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id64 == 0 {
		return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
			"id": "must be positive integer",
		})
	}

	task, err := h.taskService.GetByID(c.Request().Context(), uint(id64))
	if err != nil {
			return response.FromServiceError(c, err)
		}

	return c.JSON(http.StatusOK, dto.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Description: task.Description,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	})
}

func (h *TaskHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id64 == 0 {
		return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
			"id":"should be positive integer",
		})
	}
	
	var req dto.UpdateTaskRequest
	
	if err := c.Bind(&req); err != nil {
		return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
			"json":"invalid json",
		})
	}

	if req.Status != nil {
		s := *req.Status
		if s != types.New && s != types.InProgress && s != types.Done && s != types.Cancelled {
			return response.JSONError(c, http.StatusConflict, "validation_error", map[string]string{
				"status":"invalid status, should be 'in_progress', 'new', 'done' or 'cancelled'",
			})
		}
	}

	task, err := h.taskService.Update(c.Request().Context(), uint(id64), req.Title, req.Description, req.Status)
	if err != nil {
		/*
		if errors.Is(err, types.ErrValidation) {
			return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
				"validation":"error",
			}) // 400
		}
		if errors.Is(err, types.ErrNotFound) {
			return response.JSONError(c, http.StatusNotFound, "validation_error", map[string]string{
				"error":"task not found",
			})
		}

		return response.JSONError(c, http.StatusInternalServerError, "validation_error", map[string]string{
			"error":"internal error",
		}) // 500
		 */
		 return response.FromServiceError(c, err)
	}

	return c.JSON(http.StatusOK, dto.TaskResponse{ // 200 + DTO
		ID:        task.ID,
		Title:     task.Title,
		Description: task.Description,
		Status:      task.Status,
		UpdatedAt: task.UpdatedAt,
	})
}

func (h *TaskHandler) Delete(c echo.Context) error { // DELETE /tasks/:id
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id64 == 0 {
		return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
			"id":"should be positive integer",
		}) // 400
	}

	if err := h.taskService.Delete(c.Request().Context(), uint(id64)); err != nil {
		/*if errors.Is(err, types.ErrValidation) {
			return response.JSONError(c, http.StatusBadRequest, "validation_error", map[string]string{
				"error":"validation error",
			}) // 400
		}
		if errors.Is(err, types.ErrNotFound) {
			return response.JSONError(c, http.StatusNotFound, "validation_error", map[string]string{
				"error":"task not found",
			}) // 404
		}
		return response.JSONError(c, http.StatusInternalServerError, "validation_error", map[string]string{
			"error":"internal error",
		} ) // 500
	}*/
		return response.FromServiceError(c, err)
	}

	return c.NoContent(http.StatusNoContent) // 204 без тела
}



