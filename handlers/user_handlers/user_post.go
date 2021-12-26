package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/auth-system/database"
	"github.com/wattam/auth-system/models"
	"github.com/wattam/auth-system/services"
)

func Post(c *gin.Context) {

	user := models.User{}

	c.ShouldBindJSON(&user)

	user.Password = services.SHA256Encoder(user.Password)

	err := database.Db.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
