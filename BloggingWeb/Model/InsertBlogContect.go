package models

import (
	"gorm.io/gorm"

	view "BloggingWeb/View"
)

type Database struct {
	MainDB *gorm.DB
}

func (db Database) CreatePost(post *view.Blog) error {
	return db.MainDB.Create(post).Error
}

func (db Database) GetAllPosts(posts *[]view.Blog) error {
	return db.MainDB.Find(posts).Error
}

func (db Database) GetBlogByID(table interface{}, id int) (err error) {
	err = db.MainDB.First(table, id).Error
	return
}

func (db Database) DeleteTable(blog *view.Blog) {
	db.MainDB.Delete(&blog)
}

func (db Database) GetBlogsByUid(blog *[]view.Blog, uid int) (err error) {
	err = db.MainDB.Where("user_id = ?", uid).Find(blog).Error
	return
}

func (db Database) UpdatingBlogData(blog *view.Blog) (err error) {
	err = db.MainDB.Where("ID = ?", blog.ID).Updates(blog).Error
	return
}
