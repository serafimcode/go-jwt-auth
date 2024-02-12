package controller

import (
	"errors"
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

	rtCookie, err := c.Request.Cookie("refreshToken")

	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "Cookie not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Error retrieving cookie"})
		return
	}

	tp := model.TokensPair{RefreshToken: rtCookie.Value, AccessToken: req.AccessToken}
	newTokens, err := rtc.TokensService.RefreshTokens(tp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenExpire, err := rtc.TokensService.ExtractExpiryTime(newTokens.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    newTokens.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  refreshTokenExpire,
	}
	http.SetCookie(c.Writer, cookie)

	resp := model.TokenResponse{AccessToken: newTokens.AccessToken}
	c.JSON(http.StatusOK, resp)
}
