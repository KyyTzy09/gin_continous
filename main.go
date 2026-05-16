package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  200,
			"success": true,
			"message": "Hello Gin",
		})
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  200,
			"success": true,
			"message": "Pong",
		})
	})

	port := "8080"
	log.Printf("Starting %s server on port %s", "Gin-Continous", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf(fmt.Sprintf("Failed to start server: %v", err))
	}
}
