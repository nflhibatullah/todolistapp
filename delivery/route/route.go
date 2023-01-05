package route

import (
	"github.com/labstack/echo/v4"
	"todo/delivery/controller/activity"
	"todo/delivery/controller/todo"
)

func RegisterPath(
	e *echo.Echo, ac *activity.ActivityController, tc *todo.TodoController,
) {
	e.POST("/activity-groups", ac.CreateActivity())
	e.PATCH("/activity-groups/:id", ac.UpdateActivity())
	e.GET("/activity-groups", ac.GetAllActivity())
	e.GET("/activity-groups/:id", ac.GetActivityById())
	e.DELETE("/activity-groups/:id", ac.DeleteActivity())

	e.POST("/todo-items", tc.CreateTodo())
	e.PATCH("/todo-items/:id", tc.UpdateTodo())
	e.GET("/todo-items", tc.GetAllTodo())
	e.GET("/todo-items/:id", tc.GetTodoById())
	e.DELETE("/todo-items/:id", tc.DeleteTodo())
}
