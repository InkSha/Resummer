package auth

import "github.com/gin-gonic/gin"

func BindRouter(engine *gin.Engine) *gin.Engine {
	authGroup := engine.Group("/auth")
	authGroup.POST("/login", login)
	authGroup.POST("/register", register)
	return engine
}
