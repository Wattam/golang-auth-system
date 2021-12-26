package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

func Delete(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&models.User{}, id)

	c.Status(http.StatusNoContent)
}
