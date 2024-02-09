package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/api/controller"
	"jwt-auth/internal/domain/service"
	"jwt-auth/internal/infrastructure/repository"
	"time"
)

func NewRefreshTokenRouter(env *bootstrap.Env, duration time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewRefreshTokenRepository(db)
	ts := service.TokensService{TokenRepository: &tr, Env: env, ContextTimeout: duration}

	rtc := &controller.RefreshTokenController{TokenService: &ts}
	gtc := &controller.GetTokensController{TokensService: &ts}

	group.GET("/get-tokens", rtc.RefreshTokens)
	group.POST("/refresh", gtc.GetTokenPair)
}
