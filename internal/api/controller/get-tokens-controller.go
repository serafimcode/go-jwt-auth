package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/internal/domain/service"
)

type GetTokensController struct {
	TokensService *service.TokensService
}

func (c *GetTokensController) GetTokenPair(g *gin.Context) {

}
