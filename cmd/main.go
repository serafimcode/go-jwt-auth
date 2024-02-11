package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/api/route"
)

func main() {
	app := bootstrap.NewApp()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDbConnection()

	g := gin.Default()

	route.Init(env, db, g)

	g.Run(":" + env.ServerPort)
}
