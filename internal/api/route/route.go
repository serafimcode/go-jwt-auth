package route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth/bootstrap"
)

func Init(env *bootstrap.Env, db *mongo.Database, gin *gin.Engine) {
	router := gin.Group("/api/auth")
	NewRefreshTokenRouter(env, db, router)
}
