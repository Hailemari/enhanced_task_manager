package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Hailemari/enhanced_task_manager/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.POST("/tasks", controllers.AddTask)

	return r
}
