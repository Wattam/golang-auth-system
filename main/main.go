package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
	"github.com/wattam/user-auth-system/handlers/login_handlers"
	"github.com/wattam/user-auth-system/handlers/shoe_handlers"
	"github.com/wattam/user-auth-system/handlers/user_handlers"
	"github.com/wattam/user-auth-system/main/middlewares"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	users := r.Group("users", middlewares.Authentication())
	{
		users.GET("/get", user_handlers.GetAll)
		users.GET("/:id", user_handlers.Get)
		users.POST("/post", user_handlers.Post)
		users.PUT("/put", user_handlers.Put)
		users.DELETE("/:id", user_handlers.Delete)
	}

	shoes := r.Group("shoes", middlewares.Authentication())
	{
		shoes.GET("/get", shoe_handlers.GetAll)
		shoes.GET("/:id", shoe_handlers.Get)
		shoes.POST("/post", shoe_handlers.Post)
		shoes.PUT("/put", shoe_handlers.Put)
		shoes.DELETE("/:id", shoe_handlers.Delete)
	}

	r.POST("/login", login_handlers.Login)

	r.Run()

	defer database.DisconnectDatabase()
}
