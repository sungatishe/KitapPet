package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"unique"`
	Email          string `json:"email" gorm:"unique"`
	HashedPassword string `json:"-"`
	ReadBooks      []Book `gorm:"many2many:user_read_books;"`
	DroppedBooks   []Book `gorm:"many2many:dropped_books;"`
}

type UserReadBooks struct {
	UserID uint `gorm:"primaryKey"`
	BookID uint `gorm:"primaryKey"`
}

type UserAbandonedBooks struct {
	UserID uint `gorm:"primaryKey"`
	BookID uint `gorm:"primaryKey"`
}
