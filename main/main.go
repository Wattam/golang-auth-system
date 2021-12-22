package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
	"github.com/wattam/user-auth-system/handlers/userHandlers"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	users := r.Group("users")
	{
		users.GET("/get", userHandlers.GetAll)
		users.GET("/:id", userHandlers.Get)
		users.POST("/post", userHandlers.Post)
		users.PUT("/put", userHandlers.Put)
		users.DELETE("/:id", userHandlers.Delete)
	}

	r.Run()

	defer database.DisconnectDatabase()
}
