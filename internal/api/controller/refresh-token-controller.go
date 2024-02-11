package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/internal/domain/model"
	"jwt-auth/internal/domain/service"
	"net/http"
)

type RefreshTokenController struct {
	TokensService *service.TokensService
}

func (rtc *RefreshTokenController) RefreshTokens(c *gin.Context) {
	req := model.RefreshTokenRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	_, err := rtc.TokensService.RetrieveToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	resp, err := rtc.TokensService.RefreshAccessToken(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
