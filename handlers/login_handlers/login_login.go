package login_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/auth-system/database"
	"github.com/wattam/auth-system/models"
	"github.com/wattam/auth-system/services"
)

func Login(c *gin.Context) {

	login := models.Login{}
	c.ShouldBindJSON(&login)

	user := models.User{}

	emailError := database.Db.Where("username = ? OR email = ?", login.Credential, login.Credential).First(&user).Error
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username/email ",
		})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	token, err := services.NewJwtService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
