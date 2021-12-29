package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

// swagger:operation GET /users/get GetAll
// Get all users.
// ---
//
// responses:
//  200:
//   description: OK
//  204:
//   description: NO CONTENT
func GetAll(c *gin.Context) {

	users := []models.User{}

	database.Db.Find(&users)

	if len(users) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, users)
}
