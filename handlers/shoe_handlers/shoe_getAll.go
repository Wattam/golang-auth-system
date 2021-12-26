package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/auth-system/database"
	"github.com/wattam/auth-system/models"
)

func GetAll(c *gin.Context) {

	shoes := []models.Shoe{}

	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}