package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

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
//     style:
//      type: string
//     colour:
//      type: string
//     material:
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
func Put(c *gin.Context) {

	shoe := models.Shoe{}

	c.ShouldBindJSON(&shoe)

	if database.Db.First(&models.Shoe{}, shoe.ID).Error != nil {
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
