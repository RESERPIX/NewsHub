package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type NewsInput struct {
	Title   string `json:"title" 	binding:"required"`
	Content string `json:"content" 	binding:"required"`
}

var newsList = []News{}

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "world"})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Привет, " + name + "!"})

	})

	router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		c.JSON(http.StatusOK, gin.H{"result": "Ты искал " + query})
	})

	router.POST("/news", func(c *gin.Context) {
		var newNews News
		if err := c.ShouldBindJSON(&newNews); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		newNews.ID = len(newsList) + 1

		newsList = append(newsList, newNews)

		c.JSON(http.StatusCreated, newNews)
	})

	router.PUT("/news:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		var updatedNews News

		if err := c.ShouldBindJSON(&updatedNews); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		for i, news := range newsList {

			if news.ID == id {
				newsList[i] = updatedNews
				c.JSON(http.StatusOK, updatedNews)
				return
			}

		}
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})

	})

	router.GET("/news", func(c *gin.Context) {
		c.JSON(http.StatusOK, newsList)
	})

	router.GET("/news/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
			return
		}

		for _, news := range newsList {
			if news.ID == id {
				c.JSON(http.StatusOK, news)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
	})

	router.Run(":8081")
}
