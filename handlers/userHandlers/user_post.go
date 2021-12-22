package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
	"github.com/wattam/user-auth-system/models"
	"github.com/wattam/user-auth-system/services"
)

func Post(c *gin.Context) {

	user := models.User{}

	c.ShouldBindJSON(&user)

	user.Password = services.SHA256Encoder(user.Password)

	database.Db.Create(&user)

	c.JSON(http.StatusOK, user)
}
