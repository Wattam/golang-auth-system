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
	"github.com/wattam/golang-auth-system/controller"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/main/middleware"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("users", middleware.Authentication())
		{
			users.GET("/get", controller.GetAllUsers)
			users.GET("/:id", controller.GetUser)
			users.POST("/post", controller.PostUser)
			users.PUT("/put", controller.PutUser)
			users.DELETE("/:id", controller.DeleteUser)
		}

		shoes := v1.Group("shoes", middleware.Authentication())
		{
			shoes.GET("/get", controller.GetAllShoes)
			shoes.GET("/:id", controller.GetShoe)
			shoes.POST("/post", controller.PostShoe)
			shoes.PUT("/put", controller.PutShoe)
			shoes.DELETE("/:id", controller.DeleteShoe)
		}

		v1.POST("/login", controller.Login)
	}

	r.Run()

	defer database.DisconnectDatabase()
}
