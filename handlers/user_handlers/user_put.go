package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
	"github.com/wattam/golang-auth-system/services"
)

// swagger:operation PUT /users/put Put
// Edits a user.
// ---
//
// parameters:
// - name: user
//   in: body
//   description: The user to edit.
//   required: true
//   schema:
//    type: object
//    properties:
//     id:
//      type: number
//     username:
//      type: string
//     email:
//      type: string
//     password:
//      type: string
//
// responses:
//  200:
//   description: OK
//  204:
//   description: NO CONTENT
//  405:
//   description: METHOD NOT ALLOWED
func Put(c *gin.Context) {

	user := models.User{}

	c.ShouldBindJSON(&user)

	if database.Db.First(&models.User{}, user.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if user.Password != "" {
		user.Password = services.SHA256Encoder(user.Password)
	}

	err := database.Db.Save(&user).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
