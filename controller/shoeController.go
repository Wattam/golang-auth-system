package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/model"
)

// swagger:operation GET /shoes/get GetAll
// Get all shoes.
// ---
//
// responses:
//  200:
//   description: OK
//  204:
//   description: NO CONTENT
func GetAllShoes(c *gin.Context) {

	shoes := []model.Shoe{}

	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}

// swagger:operation GET /shoes/{ID} Get
// Gets a shoe by ID.
// ---
//
// parameters:
// - name: ID
//   in: path
//   description: ID of the shoe to get.
//   required: true
//   type: integer
//
// responses:
//  200:
//   description: OK
//  404:
//   description: NOT FOUND
func GetShoe(c *gin.Context) {

	id := c.Param("id")

	shoe := model.Shoe{}

	err := database.Db.First(&shoe, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

// swagger:operation POST /shoes/post Post
// Creates a shoe.
// ---
//
// parameters:
// - name: shoe
//   in: body
//   description: The shoe to create.
//   required: true
//   schema:
//    type: object
//    properties:
//     name:
//      type: string
//     color:
//      type: string
//     price:
//      type: number
//
// responses:
//  200:
//   description: OK
//  405:
//   description: METHOD NOT ALLOWED
func PostShoe(c *gin.Context) {

	shoe := model.Shoe{}

	c.ShouldBindJSON(&shoe)

	err := database.Db.Create(&shoe).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

// swagger:operation PUT /shoes/put Put
// Edits a shoe.
// ---
//
// parameters:
// - name: shoe
//   in: body
//   description: The shoe to edit.
//   required: true
//   schema:
//    type: object
//    properties:
//     id:
//      type: number
//     name:
//      type: string
//     colorr:
//      type: string
//     price:
//      type: number
//
// responses:
//  200:
//   description: OK
//  204:
//   description: NO CONTENT
//  405:
//   description: METHOD NOT ALLOWED
func PutShoe(c *gin.Context) {

	shoe := model.Shoe{}

	c.ShouldBindJSON(&shoe)

	if database.Db.First(&model.Shoe{}, shoe.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err := database.Db.Save(&shoe).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

// swagger:operation DELETE /shoes/{ID} Delete
// Deletes a shoe by ID.
// ---
//
// parameters:
// - name: ID
//   in: path
//   description: ID of the shoe to delete.
//   required: true
//   type: integer
//
// responses:
//  204:
//   description: NO CONTENT
func DeleteShoe(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&model.Shoe{}, id)

	c.Status(http.StatusNoContent)
}
