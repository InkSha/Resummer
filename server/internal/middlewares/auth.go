package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// authorization := c.GetHeader("Authorization")
		// if authorization == "" {
		// 	c.JSON(http.StatusUnauthorized, network.Failed[*string]("Authorization header is required", nil))
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
