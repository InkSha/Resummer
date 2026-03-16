package auth

import (
	"fmt"
	"net/http"

	"github.com/InkSha/Resummer/internal/config"
	"github.com/InkSha/Resummer/internal/service/repository"
	"github.com/InkSha/Resummer/pkg/network"
	"github.com/gin-gonic/gin"
)

// @BasePath	/auth
// @Summary	Register a new user
// @Schemes
// @Description	Register a new user with account and password.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		AuthRequestDTO	true	"Register Request"
// @Success 200 {object} network.Response
// @Failure 400 {object} network.Response
// @Failure 409 {object} network.Response
// @Failure 500 {object} network.Response
// @Router			/register [post]
func register(c *gin.Context) {
	var dto AuthRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, network.Failed(err.Error(), nil))
		return
	}

	user, err := repository.UserRegister(dto.Account, dto.Password)
	if err != nil {
		c.JSON(http.StatusConflict, network.Failed(err.Error(), nil))
		return
	}

	token, err := GenToken(fmt.Sprint(user.ID), config.Get("JWT_SECRET"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, network.Failed(err.Error(), nil))
		return
	}

	// c.JSON(http.StatusOK, network.Success(token))
	c.JSON(http.StatusOK, network.Success(token))
}
