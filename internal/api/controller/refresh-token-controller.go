package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/internal/domain/service"
)

type RefreshTokenController struct {
	TokenService *service.TokensService
}

func (rtc *RefreshTokenController) RefreshTokens(c *gin.Context) {

}
