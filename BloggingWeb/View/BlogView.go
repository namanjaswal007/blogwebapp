package view

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	ID      int    `json:"ID" gorm:"primarykey" form:"ID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}
