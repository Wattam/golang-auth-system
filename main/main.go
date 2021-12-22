package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	users := r.Group("users")
	{
		users.GET("")
		users.GET("")
		users.POST("")
		users.PUT("")
		users.DELETE("")
	}

	r.Run()

	defer database.DisconnectDatabase()
}
