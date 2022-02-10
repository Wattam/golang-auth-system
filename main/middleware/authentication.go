package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/service"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		const bearerSchema = "Bearer "

		header := c.GetHeader("Authorization")

		if header == "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := header[len(bearerSchema):]

		if !service.NewJwtService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
