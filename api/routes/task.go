package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// TaskRoutes struct
type TaskRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	taskController controllers.TaskController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup task routes
func (s TaskRoutes) Setup() {
	s.logger.Info("Setting up task routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/task", s.taskController.GetTask)
		api.GET("/task/:id", s.taskController.GetOneTask)
		api.POST("/task", s.taskController.SaveTask)
		api.POST("/task/:id", s.taskController.UpdateTask)
		api.DELETE("/task/:id", s.taskController.DeleteTask)
	}
}

// NewUserRoutes creates new user controller
func NewTaskRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	taskController controllers.TaskController,
	authMiddleware middlewares.JWTAuthMiddleware,
) TaskRoutes {
	return TaskRoutes{
		handler:        handler,
		logger:         logger,
		taskController: taskController,
		authMiddleware: authMiddleware,
	}
}
