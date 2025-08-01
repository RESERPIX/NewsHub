package main

import (
	"news-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/news", handlers.CreateNewsHandler)
	router.GET("/news", handlers.GetAllNewsHandler)
	router.GET("/news/:id", handlers.GetNewsByIDHandler)
	router.PUT("/news/:id", handlers.UpdateNewsHandler)
	router.DELETE("/news/:id", handlers.DeleteNewsHandler)

	router.Run(":8081")
}
