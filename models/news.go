package models

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"title" binding:"required"`
}

var NewsList []News
