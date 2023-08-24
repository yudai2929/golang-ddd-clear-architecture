package controllers

import (
	"fmt"
	"golang-ddd-clear-architecture/day4/task3/adapter/messages"
	"golang-ddd-clear-architecture/day4/task3/adapter/requests"
	"golang-ddd-clear-architecture/day4/task3/usecase"
	"golang-ddd-clear-architecture/day4/task3/usecase/params"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	tu usecase.TaskUsecase
}

func NewTaskController(TaskUsecase usecase.TaskUsecase) TaskController {
	return TaskController{tu: TaskUsecase}
}

func (tc TaskController) Index(c echo.Context) error {
	tasks, err := tc.tu.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc TaskController) Show(c echo.Context) error {
	intID, _ := strconv.Atoi(c.Param("id"))
	task, err := tc.tu.FindByID(intID)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, task)
}

func (tc TaskController) Create(c echo.Context) error {

	req := new(requests.TaskRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	params := params.TaskParams{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
	}

	if err := tc.tu.Create(params); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("Create Task Success"))

}

func (tc TaskController) Update(c echo.Context) error {
	req := new(requests.TaskRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	params := params.TaskParams{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
	}

	if err := tc.tu.Update(taskID, params); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("Update Task Success"))
}

func (tc TaskController) Delete(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := tc.tu.Delete(taskID); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("Delete Task Success"))
}

func (tc TaskController) Complete(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := tc.tu.Complete(taskID); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("Complete Task Success"))
}

func (tc TaskController) OnHold(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := tc.tu.OnHold(taskID); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("OnHold Task Success"))
}

func (tc TaskController) InProgress(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := tc.tu.InProgress(taskID); err != nil {
		return err
	}

	return c.JSON(200, messages.SuccessMessage("InProgress Task Success"))
}
