package view

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int    `gorm:"unique" json:"ID"`
	Email     string `gorm:"unique" json:"email"`
	FullName  string `json:"full_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type UserSessionToken struct {
	Username string    `json:"user_name"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
	Exp      time.Time `json:"exp"`
}

type UserCredentials struct {
	ID        int    `json:"ID"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type UserSession struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	UserAgent string    `json:"user_agent"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
