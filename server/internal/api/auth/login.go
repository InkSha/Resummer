package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/InkSha/Resummer/internal/config"
	"github.com/InkSha/Resummer/internal/service/repository"
	"github.com/InkSha/Resummer/pkg/network"
	"github.com/gin-gonic/gin"
)

// @BasePath /auth
// @Summary User login
// @Schemes
// @Description User login
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		AuthRequestDTO	true	"Login Request"
// @Success 200 {object} network.Response
// @Failure 400 {object} network.Response
// @Failure 401 {object} network.Response
// @Failure 500 {object} network.Response
// @Router			/login [post]
func login(c *gin.Context) {
	var dto AuthRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, network.Failed(err.Error(), nil))
		return
	}

	// login
	user, err := repository.UserLogin(dto.Account, dto.Password)
	if err != nil {
		if errors.Is(err, repository.ErrPasswordIncorrect) {
			c.JSON(http.StatusUnauthorized, network.Failed(err.Error(), nil))
			return
		}
		if errors.Is(err, repository.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, network.Failed(err.Error(), nil))
			return
		}
		c.JSON(http.StatusBadRequest, network.Failed(err.Error(), nil))
		return
	}

	// generate token
	token, err := GenToken(fmt.Sprint(user.ID), config.Get("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, network.Failed(err.Error(), nil))
		return
	}

	// return token
	c.JSON(http.StatusOK, network.Success(token))
}
