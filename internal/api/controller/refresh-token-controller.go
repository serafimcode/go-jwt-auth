package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/internal/domain/service"
)

type RefreshTokenController struct {
	TokenService *service.TokensService
}

func (c *RefreshTokenController) RefreshTokens(g *gin.Context) {

}
