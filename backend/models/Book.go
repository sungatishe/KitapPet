package models

import "gorm.io/gorm"

type BookType string

const (
	Manga  BookType = "Манга"
	Comics BookType = "Комиксы"
	Books  BookType = "Книги"
)

type Book struct {
	gorm.Model
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Cover       string   `json:"cover"`
	Description string   `json:"description"`
	Year        string   `json:"year"`
	Type        BookType `json:"type"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Type        string `json:"type"`
}
