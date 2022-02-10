package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/model"
	"github.com/wattam/golang-auth-system/service"
)

// swagger:operation POST /login Login
// Logins a user.
// ---
//
// parameters:
// - name: login
//   in: body
//   description: The user credentials to login.
//   required: true
//   schema:
//    type: object
//    properties:
//     credential:
//      type: string
//     password:
//      type: string
//
// responses:
//  200:
//   description: OK
//  400:
//   description: BAD REQUEST
//  401:
//   description: UNAUTHORIZED
//  500:
//   description: INTERNAL SERVER ERROR
func Login(c *gin.Context) {

	login := model.Login{}
	c.ShouldBindJSON(&login)

	user := model.User{}

	emailError := database.Db.Where("username = ? OR email = ?", login.Credential, login.Credential).First(&user).Error
	if emailError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username/email ",
		})
		return
	}

	if user.Password != service.SHA256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	token, err := service.NewJwtService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
