package main

import (
	"github.com/InkSha/Resummer/auth"
	"github.com/InkSha/Resummer/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /
// @Summary Index
// @Schemes
// @Description This is a Resummer API server root api.
// @Tags Root, Example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", index)

	auth.BindRouter(router.Group("/auth"))

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
