package main

import (
	"awesomeProject/db"
	"awesomeProject/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	db.InitRedis()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:8090"}, // разрешите нужные домены
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("", services.CreateBook)
		bookRoutes.POST("/all", services.CreateBooks)
		bookRoutes.GET("/:id", services.GetBookById)
		bookRoutes.GET("", services.GetAllBooks)
		bookRoutes.PUT("/:id", services.UpdateBook)
		bookRoutes.DELETE("/:id", services.DeleteBook)
	}

	r.GET("/ws", func(c *gin.Context) {
		services.HandleConnection(c.Writer, c.Request)
	})

	// Запуск обработчика WebSocket-сообщений в отдельной горутине
	go services.HandleMessages()

	r.POST("/register", services.Register)
	r.POST("/login", services.Login)
	r.GET("/users/me", services.GetUserData)

	userRoutes := r.Group("/users/:user_id")
	userRoutes.Use(services.AuthMiddleware())
	{
		userRoutes.POST("/read-books/:book_id", services.AddReadBook)
		userRoutes.DELETE("/read-books/:book_id", services.RemoveReadBook)
		userRoutes.GET("/read-books", services.GetReadBooks)

		userRoutes.POST("/abandoned-books/:book_id", services.AddDroppedBook)
		userRoutes.DELETE("/abandoned-books/:book_id", services.RemoveDroppedBook)
		userRoutes.GET("/abandoned-books", services.GetDroppedBooks)
	}

	r.Run(":8080")
}
