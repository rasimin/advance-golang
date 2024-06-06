package main

import (
	"advance/middleware"
	"advance/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	// Definisikan route
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Halo dari Gin!",
	// 	})
	// })

	// r.POST("/post", func(c *gin.Context) {
	// 	var json struct {
	// 		Message string `json:"message"`
	// 	}
	// 	if err := c.ShouldBindJSON(&json); err == nil {
	// 		c.JSON(200, gin.H{"message": json.Message})
	// 	} else {
	// 		c.JSON(400, gin.H{"error": err.Error()})
	// 	}
	// })

	router.SetupRouter(r)
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
