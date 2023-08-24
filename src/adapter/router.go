package adapter

import (
	"golang-ddd-clear-architecture/day3/mysql"
	"golang-ddd-clear-architecture/day4/task3/adapter/controllers"
	"golang-ddd-clear-architecture/day4/task3/infrastructure"
	"golang-ddd-clear-architecture/day4/task3/usecase"

	"github.com/labstack/echo/v4"
)

func InitRouter() {
	e := echo.New()

	db, err := mysql.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskRepository := infrastructure.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controllers.NewTaskController(taskUsecase)

	e.GET("/tasks", taskController.Index)
	e.GET("/tasks/:id", taskController.Show)
	e.POST("/tasks", taskController.Create)
	e.PUT("/tasks/:id", taskController.Update)
	e.DELETE("/tasks/:id", taskController.Delete)
	e.PUT("/tasks/:id/complete", taskController.Complete)
	e.PUT("/tasks/:id/onload", taskController.OnHold)
	e.PUT("/tasks/:id/inProgress", taskController.InProgress)

	e.Logger.Fatal(e.Start(":8080"))
}
