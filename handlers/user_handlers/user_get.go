package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

// swagger:operation GET /users/{ID} Get
// Gets a user by ID.
// ---
//
// parameters:
// - name: ID
//   in: path
//   description: ID of the user to get.
//   required: true
//   type: integer
//
// responses:
//  200:
//   description: OK
//  404:
//   description: NOT FOUND
func Get(c *gin.Context) {

	user := models.User{}
	id := c.Param("id")

	err := database.Db.First(&user, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
