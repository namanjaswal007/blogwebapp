package view

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `gorm:"unique" json:"ID"`
	Email string `gorm:"unique" json:"email"`
	Name  string `json:"name"`
}
