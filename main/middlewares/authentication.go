package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/user-auth-system/services"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		const bearerSchema = "Bearer "

		header := c.GetHeader("Authorization")

		//fmt.Printf("\n\n%v\n\n", header)

		if header == "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := header[len(bearerSchema):]

		if !services.NewJwtService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
