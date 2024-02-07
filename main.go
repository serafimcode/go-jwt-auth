package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt-auth/database"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	dbClient := database.InitDb()
	fmt.Println(dbClient)

	if port == "" {
		port = "8000"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	router.Run(":" + port)
}
