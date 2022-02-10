package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/model"
	"github.com/wattam/golang-auth-system/service"
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
func GetAllUsers(c *gin.Context) {

	users := []model.User{}

	database.Db.Find(&users)

	if len(users) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, users)
}

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
func GetUser(c *gin.Context) {

	user := model.User{}
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

// swagger:operation POST /users/post Post
// Creates a user.
// ---
//
// parameters:
// - name: user
//   in: body
//   description: The user to create.
//   required: true
//   schema:
//    type: object
//    properties:
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
//  405:
//   description: METHOD NOT ALLOWED
func PostUser(c *gin.Context) {

	user := model.User{}

	c.ShouldBindJSON(&user)

	if user.Password != "" {
		user.Password = service.SHA256Encoder(user.Password)
	}

	err := database.Db.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

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
func PutUser(c *gin.Context) {

	user := model.User{}

	c.ShouldBindJSON(&user)

	if database.Db.First(&model.User{}, user.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	if user.Password != "" {
		user.Password = service.SHA256Encoder(user.Password)
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

// swagger:operation DELETE /users/{ID} Delete
// Deletes a user by ID.
// ---
//
// parameters:
// - name: ID
//   in: path
//   description: ID of the user to delete.
//   required: true
//   type: integer
//
// responses:
//  204:
//   description: NO CONTENT
func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&model.User{}, id)

	c.Status(http.StatusNoContent)
}
