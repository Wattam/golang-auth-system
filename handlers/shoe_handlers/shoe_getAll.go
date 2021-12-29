package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
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
func GetAll(c *gin.Context) {

	shoes := []models.Shoe{}

	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}
