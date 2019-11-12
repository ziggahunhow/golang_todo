package routes

import (
	"Todo/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("Todo/:id", controllers.GetATodo)
		v1.PUT("Todo/:id", controllers.UpdateATodo)
		v1.DELETE("Todo/:id", controllers.DeleteATodo)
	}
	return r
}
