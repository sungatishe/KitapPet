package services

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Db.Create(&book)
	c.JSON(http.StatusOK, book)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.Db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	fmt.Println(db.Db.Find(&books))
	cachedBooks, err := db.RedisClient.Get(db.Ctx, "books_cache").Result()
	if err == redis.Nil {
		db.Db.Find(&books)
		bookJson, err := json.Marshal(books)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сериализации данных"})
			return
		}
		err = db.RedisClient.Set(db.Ctx, "books_cache", bookJson, 10*time.Minute).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сериализации данных"})
			return
		}
		log.Println("Books cached successfully")
		c.JSON(http.StatusOK, books)
	} else if err != nil {
		log.Println("Error retrieving data from Redis:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка Redis"})
		return
	} else {
		// Если книги найдены в кэше, десериализуем их
		log.Println("Books retrieved from cache")
		err = json.Unmarshal([]byte(cachedBooks), &books)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка десериализации данных"})
			return
		}

		// Возвращаем книги из кэша
		c.JSON(http.StatusOK, books)
	}

}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.Db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	db.Db.Save(book)
	c.JSON(http.StatusOK, book)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.Db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func CreateBooks(c *gin.Context) {
	var booksInput []models.CreateBookInput

	// Привязка входных данных из JSON к структуре
	if err := c.ShouldBindJSON(&booksInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Инициализируем пустой список для книг
	var books []models.Book

	// Цикл по входным данным
	for _, input := range booksInput {
		book := models.Book{
			Title:       input.Title,
			Author:      input.Author,
			Cover:       input.Cover,
			Description: input.Description,
			Year:        input.Year,
			Type:        models.BookType(input.Type),
		}

		books = append(books, book)
	}

	if err := db.Db.Create(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
