package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/database"
	"github.com/wattam/user-auth-system/models"
	"github.com/wattam/user-auth-system/services"
)

func Put(c *gin.Context) {

	user := models.User{}

	c.ShouldBindJSON(&user)

	if database.Db.First(&models.User{}, user.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	database.Db.Save(&user)

	c.JSON(http.StatusOK, user)
}
