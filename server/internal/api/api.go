package api

import (
	"strconv"

	"github.com/InkSha/Resummer/cmd/server/docs"
	"github.com/InkSha/Resummer/internal/api/auth"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath	/
// @Summary	Index
// @Schemes
// @Description	This is a Resummer API server root api.
// @Tags			Root, Example
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]string
// @Router			/ [get]
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func Listen(port int) {
	router := gin.Default()

	router.GET("/", index)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth.BindRouter(router)

	router.Run("localhost:" + strconv.Itoa(port))
}
