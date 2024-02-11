package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/api/controller"
	"jwt-auth/internal/domain/service"
	"jwt-auth/internal/infrastructure/repository"
)

func NewRefreshTokenRouter(env *bootstrap.Env, db *mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewRefreshTokenRepository(db)
	ts := service.TokensService{TokenRepository: tr, Env: env}

	gtc := &controller.GetTokensController{TokensService: &ts}
	rtc := &controller.RefreshTokenController{TokensService: &ts}

	group.POST("/get-tokens", gtc.GetTokens)
	group.POST("/refresh", rtc.RefreshTokens)
}
