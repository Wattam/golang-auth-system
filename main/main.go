package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
	"github.com/wattam/user-auth-system/handlers/user_handlers"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	users := r.Group("users")
	{
		users.GET("/get", user_handlers.GetAll)
		users.GET("/:id", user_handlers.Get)
		users.POST("/post", user_handlers.Post)
		users.PUT("/put", user_handlers.Put)
		users.DELETE("/:id", user_handlers.Delete)
	}

	r.Run()

	defer database.DisconnectDatabase()
}
