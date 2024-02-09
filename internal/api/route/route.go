package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/bootstrap"
	"time"
)

func Init(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	router := gin.Group("/auth")
	NewRefreshTokenRouter(env, timeout, db, router)
}
