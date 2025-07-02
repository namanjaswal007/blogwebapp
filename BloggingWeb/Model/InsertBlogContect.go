package models

import (
	view "BloggingWeb/View"

	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, post *view.Post) error {
	return db.Create(post).Error
}

func GetAllPosts(db *gorm.DB, posts *[]view.Post) error {
	return db.Find(posts).Error
}
