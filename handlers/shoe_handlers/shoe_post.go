package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

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
//     style:
//      type: string
//     colour:
//      type: string
//     material:
//      type: string
//     price:
//      type: "number"
//
// responses:
//  200:
//   description: OK
//  405:
//   description: METHOD NOT ALLOWED
func Post(c *gin.Context) {

	shoe := models.Shoe{}

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
