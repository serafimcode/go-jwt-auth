package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/api/route"
	"time"
)

func main() {
	app := bootstrap.NewApp()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDbConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	g := gin.Default()

	route.Init(env, timeout, db, g)

	g.Run(":" + env.ServerPort)
}
