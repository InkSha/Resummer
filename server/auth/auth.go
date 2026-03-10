package auth

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login success",
	})
}

func register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register success",
	})
}

func BindRouter(group *gin.RouterGroup) {
	group.POST("/login", login)
	group.POST("/register", register)
}
