package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt-auth/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	fmt.Println(db)
	defer app.CloseDbConnection()

	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	g.Run(":" + env.ServerPort)
}
