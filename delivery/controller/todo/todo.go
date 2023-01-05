package todo

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"todo/delivery/common"
	"todo/entity"
	"todo/repository/todo"
)

type TodoController struct {
	TodoRepo todo.TodoRepository
}

func NewTodoController(todotRepo todo.TodoRepository) *TodoController {
	return &TodoController{todotRepo}
}

func (ac TodoController) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &CreateToDoRequest{}

		c.Bind(&data)

		if data.Title == "" {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Bad Request", "title cannot be null"))
		}
		if data.ActivityGroupID < 1 {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Bad Request", "activity_group_id cannot be null"))
		}

		var todoData = &entity.ToDo{
			ActivityGroupID: data.ActivityGroupID,
			Title:           data.Title,
			IsActive:        data.IsActive,
		}

		result, err := ac.TodoRepo.CreateTodo(todoData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Create Todo", err.Error()))
		}

		response := &DataResponse{
			Id:              result.ID,
			Title:           result.Title,
			ActivityGroupID: result.ActivityGroupID,
			IsActive:        result.IsActive,
			Priority:        result.Priority,
			CreatedAt:       result.CreatedAt,
			UpdatedAt:       result.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, common.SuccessResponse(response))
	}
}

func (ac TodoController) GetTodoById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Id must be a number", err.Error()))
		}

		result, err := ac.TodoRepo.GetTodoByID(int64(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, common.ToDoNotFoundResponse(uint(id)))
		}

		response := &DataResponse{
			Id:              result.ID,
			Title:           result.Title,
			ActivityGroupID: result.ActivityGroupID,
			IsActive:        result.IsActive,
			Priority:        result.Priority,
			CreatedAt:       result.CreatedAt,
			UpdatedAt:       result.UpdatedAt,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac TodoController) GetAllTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			groupID int
			query   = c.QueryParam("activity_group_id")
			err     error
		)

		groupID, err = strconv.Atoi(query)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("activity_group_id cannot be null", err.Error()))
		}

		result, err := ac.TodoRepo.GetAllTodo(int64(groupID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Get All Todo", err.Error()))
		}

		var response = make([]*DataResponse, 0)
		for _, v := range result {
			response = append(response, &DataResponse{
				Id:              v.ID,
				Title:           v.Title,
				ActivityGroupID: v.ActivityGroupID,
				IsActive:        v.IsActive,
				Priority:        v.Priority,
				CreatedAt:       v.CreatedAt,
				UpdatedAt:       v.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac TodoController) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &UpdateToDoRequest{}

		c.Bind(&data)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Id must be a number", err.Error()))
		}

		var todo = &entity.ToDo{Title: data.Title, Priority: data.Priority, IsActive: data.IsActive, ActivityGroupID: data.ActivityGroupID}
		todo.ID = uint(id)
		todo.UpdatedAt = time.Now()

		result, err := ac.TodoRepo.UpdateTodo(todo)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, common.ToDoNotFoundResponse(uint(id)))
			}
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Update Todo", err.Error()))
		}

		response := &DataResponse{
			Id:              result.ID,
			Title:           result.Title,
			ActivityGroupID: result.ActivityGroupID,
			IsActive:        result.IsActive,
			Priority:        result.Priority,
			CreatedAt:       result.CreatedAt,
			UpdatedAt:       result.UpdatedAt,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(response))
	}
}

func (ac TodoController) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse("Id must be a number", err.Error()))
		}

		err = ac.TodoRepo.DeleteTodoByID(int64(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, common.ToDoNotFoundResponse(uint(id)))
			}
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse("Failed Delete Todo", err.Error()))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(make(map[string]interface{})))
	}
}
