package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"todo/config"
	activityCtrl "todo/delivery/controller/activity"
	todoCtrl "todo/delivery/controller/todo"
	"todo/delivery/route"
	activityRepo "todo/repository/activity"
	todoRepo "todo/repository/todo"
	"todo/util"
)

func main() {

	config := config.GetConfig()
	db := util.InitDB(config)
	e := echo.New()

	activityRepo := activityRepo.NewActivityRepository(db)
	activityController := activityCtrl.NewActivityController(activityRepo)

	todoRepo := todoRepo.NewTodoRepository(db)
	todoController := todoCtrl.NewTodoController(todoRepo)

	route.RegisterPath(e, activityController, todoController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", 3030)))
}
