package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-dynamodb-example/handlers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/users")
	{
		api.POST("/", handlers.CreateUser)
		api.GET("/", handlers.GetUsers)
		api.GET("/:id", handlers.GetUserByID)
		api.PUT("/:id", handlers.UpdateUser)
		api.DELETE("/:id", handlers.DeleteUser)
	}

	return r
}
