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

	resp, err := gtc.TokensService.CreateTokens(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	err = gtc.TokensService.SaveRefreshToken(req.Guid, resp.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
