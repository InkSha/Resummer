package api

import (
	"strconv"

	"github.com/InkSha/Resummer/cmd/server/docs"
	"github.com/InkSha/Resummer/internal/api/auth"
	"github.com/InkSha/Resummer/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Listen(prefix string, port int) {
	service := gin.Default()

	router := service.Group(prefix)
	auth.BindRouter(router)

	internal := router.Group(prefix)
	internal.Use(middlewares.AuthMiddleware())

	docs.SwaggerInfo.BasePath = prefix
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	service.Run("localhost:" + strconv.Itoa(port))
}
