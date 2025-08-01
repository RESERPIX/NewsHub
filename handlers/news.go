package handlers

import (
	"NewsHub/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewsHandler(c *gin.Context) {
	var input models.NewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	newNews := models.News{
		ID:      len(models.Newslist) + 1,
		Title:   input.Title,
		Content: input.Content,
	}
	models.NewsList = append(models.NewsList, newNews)
	c.JSON(http.StatusCreated, newNews)
}

func GetNewsByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}

	for _, news := range models.NewsList {
		if news.ID == id {
			c.JSON(http.StatusOK, news)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"}) // Добавлено: обработка случая "не найдено"
}

func GetAllNewsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewsList)
}

func UpdateNewsHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"}) // Уточнено сообщение об ошибке
		return
	}

	var input models.NewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"}) // Исправлено: gin.H вместо c.JSON
		return
	}

	for i, news := range models.NewsList { // Исправлено: models.NewsList вместо NewsList
		if news.ID == id {
			models.NewsList[i].Title = input.Title
			models.NewsList[i].Content = input.Content
			c.JSON(http.StatusOK, models.NewsList[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Новость не найдена"})
}

func DeleteNewsHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}

	for news, id := range NewsList {
		if id == news.ID {
			NewsList = append(NewsList[:i], NewsList[i+1:]...)
		}
	}

}
