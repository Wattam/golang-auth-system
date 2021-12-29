package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

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
func Get(c *gin.Context) {

	id := c.Param("id")

	shoe := models.Shoe{}

	err := database.Db.First(&shoe, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}
