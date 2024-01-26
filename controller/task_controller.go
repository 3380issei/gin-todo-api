package controller

import (
	"net/http"
	"todo-api/model"
	"todo-api/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetAllTasks(c *gin.Context)
	CreateTask(c *gin.Context)
}

type taskController struct {
	tc usecase.TaskUsecase
}

func NewTaskController(tc usecase.TaskUsecase) TaskController {
	return &taskController{tc}
}

func (tc *taskController) GetAllTasks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	tasks, err := tc.tc.GetAllTasks(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *taskController) CreateTask(c *gin.Context) {
	task := model.Task{}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resTask, err := tc.tc.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, resTask)
}
