package services

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddReadBook - добавление книги в список прочитанных
func AddReadBook(c *gin.Context) {
	// Извлекаем userID из контекста
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	bookID := c.Param("book_id")

	var user models.User
	if err := db.Db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var book models.Book
	if err := db.Db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Добавляем книгу в список прочитанных
	err := db.Db.Model(&user).Association("ReadBooks").Append(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Read book added"})
	broadcast <- Message{
		Action: "add",
		BookID: book.ID,
	}
}

// RemoveReadBook - удаление книги из списка прочитанных
func RemoveReadBook(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	bookID := c.Param("book_id")

	var user models.User
	if err := db.Db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var book models.Book
	if err := db.Db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Удаляем книгу из списка прочитанных
	err := db.Db.Model(&user).Association("ReadBooks").Delete(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Read book removed"})
	broadcast <- Message{
		Action: "remove",
		BookID: book.ID,
	}
}

// GetReadBooks - получение списка прочитанных книг
func GetReadBooks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.Db.Preload("ReadBooks").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user.ReadBooks)
}

// AddDroppedBook - добавление книги в список заброшенных
func AddDroppedBook(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	bookID := c.Param("book_id")

	var user models.User
	if err := db.Db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var book models.Book
	if err := db.Db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Добавляем книгу в список заброшенных
	err := db.Db.Model(&user).Association("DroppedBooks").Append(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dropped book added"})
	broadcast <- Message{
		Action: "drop",
		BookID: book.ID,
	}
}

// RemoveDroppedBook - удаление книги из списка заброшенных
func RemoveDroppedBook(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	bookID := c.Param("book_id")

	var user models.User
	if err := db.Db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var book models.Book
	if err := db.Db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Удаляем книгу из списка заброшенных
	err := db.Db.Model(&user).Association("DroppedBooks").Delete(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dropped book removed"})
	broadcast <- Message{
		Action: "remove",
		BookID: book.ID,
	}
}

// GetDroppedBooks - получение списка заброшенных книг
func GetDroppedBooks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.Db.Preload("DroppedBooks").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user.DroppedBooks)
}
