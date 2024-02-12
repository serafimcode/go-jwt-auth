package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/internal/domain/model"
	"jwt-auth/internal/domain/service"
	"net/http"
)

type GetTokensController struct {
	TokensService *service.TokensService
}

func (gtc *GetTokensController) GetTokens(c *gin.Context) {
	req := model.GetTokenRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	tokenPair, err := gtc.TokensService.CreateTokens(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenExpire, err := gtc.TokensService.ExtractExpiryTime(tokenPair.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    tokenPair.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  refreshTokenExpire,
	}
	http.SetCookie(c.Writer, cookie)

	resp := model.TokenResponse{AccessToken: tokenPair.AccessToken}
	c.JSON(http.StatusOK, resp)
}
