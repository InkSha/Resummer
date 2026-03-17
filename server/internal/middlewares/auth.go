package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/InkSha/Resummer/internal/config"
	"github.com/InkSha/Resummer/internal/globals"
	"github.com/InkSha/Resummer/pkg/network"
	"github.com/InkSha/Resummer/pkg/token"
	"github.com/gin-gonic/gin"
)

const bearerPrefix = "Bearer "

func AuthMiddleware() gin.HandlerFunc {
	jwtSecret := config.Get("JWT_SECRET")

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, network.Failed(token.ErrTokenInvalidOrExpired.Error(), nil))
			c.Abort()
			return
		}

		if !strings.HasPrefix(authorization, bearerPrefix) {
			c.JSON(http.StatusUnauthorized, network.Failed(token.ErrTokenInvalidOrExpired.Error(), nil))
			c.Abort()
			return
		}

		authorization = authorization[len(bearerPrefix):]

		if info, err := token.CheckToken(authorization, jwtSecret); err != nil {
			c.JSON(http.StatusUnauthorized, network.Failed(err.Error(), nil))
			c.Abort()
			return
		} else {
			if id, ok := info[token.UserIDKey]; ok {
				c.Set(globals.ContextUserIdKey, fmt.Sprint(id))
			}
		}

		c.Next()
	}
}
