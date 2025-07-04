package view

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id    int    `gorm:"unique" json:"ID"`
	Email string `gorm:"unique" json:"email"`
	Name  string `json:"name"`
}

type RequestData struct {
	Username string    `json:"user_name"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
	Uid      int       `json:"uid"`
	Password string    `json:"password"`
	Exp      time.Time `json:"exp"`
}

// type CustomPayload struct {
// 	RequestData RequestData
// 	Exp         time.Time `json:"exp"`
// }
