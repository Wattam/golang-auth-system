package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/auth-system/database"
	"github.com/wattam/auth-system/models"
)

func GetAll(c *gin.Context) {

	users := []models.User{}

	database.Db.Find(&users)

	if len(users) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, users)
}
