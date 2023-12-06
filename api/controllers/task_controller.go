package controllers

import (
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController data type
type TaskController struct {
	service domains.TaskService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewTaskController(taskService domains.TaskService, logger lib.Logger) TaskController {
	return TaskController{
		service: taskService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (t TaskController) GetOneTask(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	task, err := t.service.GetOneTask(uint(id))

	if err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": task,
	})

}

// GetUser gets the user
func (t TaskController) GetTask(c *gin.Context) {
	tasks, err := t.service.GetAllTask()
	if err != nil {
		t.logger.Error(err)
	}
	c.JSON(200, gin.H{"data": tasks})
}

func (t TaskController) SaveTask(c *gin.Context) {
	task := models.Task{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&task); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.service.WithTrx(trxHandle).CreateTask(task); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "task created"})
}

// UpdateUser updates user
func (t TaskController) UpdateTask(c *gin.Context) {
	task := models.Task{}

	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.service.WithTrx(trxHandle).UpdateTask(uint(id), task); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"data": "task updated"})
}

// DeleteUser deletes user
func (t TaskController) DeleteTask(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := t.service.DeleteTask(uint(id)); err != nil {
		t.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "task deleted"})
}
