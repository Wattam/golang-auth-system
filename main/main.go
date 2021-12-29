// golang-auth-system
//
// REST API, with authentication, to manage user and shoes.
//
// Schemes: [http, https]
// Host: localhost:8080
// BasePath: /api/v1
// Version: 1.0
//
// Consumes:
//  - application/json
//
// Produces:
//  - application/json
// swagger:meta

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/handlers/login_handlers"
	"github.com/wattam/golang-auth-system/handlers/shoe_handlers"
	"github.com/wattam/golang-auth-system/handlers/user_handlers"
	"github.com/wattam/golang-auth-system/main/middlewares"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("users", middlewares.Authentication())
		{
			users.GET("/get", user_handlers.GetAll)
			users.GET("/:id", user_handlers.Get)
			users.POST("/post", user_handlers.Post)
			users.PUT("/put", user_handlers.Put)
			users.DELETE("/:id", user_handlers.Delete)
		}

		shoes := v1.Group("shoes", middlewares.Authentication())
		{
			shoes.GET("/get", shoe_handlers.GetAll)
			shoes.GET("/:id", shoe_handlers.Get)
			shoes.POST("/post", shoe_handlers.Post)
			shoes.PUT("/put", shoe_handlers.Put)
			shoes.DELETE("/:id", shoe_handlers.Delete)
		}

		v1.POST("/login", login_handlers.Login)
	}

	r.Run()

	defer database.DisconnectDatabase()
}
