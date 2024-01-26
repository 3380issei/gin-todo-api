package router

import (
	"todo-api/controller"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.UserController, tc controller.TaskController, am middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.Use(am.CORS())

	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)

	t := r.Group("/tasks")
	t.Use(am.JWTAuth)
	{
		t.GET("", tc.GetAllTasks)
		t.POST("", tc.CreateTask)
		t.DELETE("/:id", tc.DeleteTask)
	}

	return r
}
