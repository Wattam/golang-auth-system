package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

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
func Delete(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&models.Shoe{}, id)

	c.Status(http.StatusNoContent)
}
