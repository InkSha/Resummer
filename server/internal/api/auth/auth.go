package auth

import "github.com/gin-gonic/gin"

func BindRouter(group *gin.RouterGroup) *gin.RouterGroup {
	api := group.Group("/auth")
	{
		api.POST("/login", login)
		api.POST("/register", register)
	}
	return api
}
