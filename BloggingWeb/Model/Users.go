package models

import (
	view "BloggingWeb/View"

	"gorm.io/gorm"
)

func (db Database) GetAllUsers(user *[]view.User) error {
	return db.MainDB.Find(user).Error
}

func (db Database) GetUserByUid(user interface{}, uid int) (err error) {
	err = db.MainDB.First(user, uid).Error
	return
}

func (db Database) AddUserDetails(user *view.User) {
	db.MainDB.Create(user)

}

func (db Database) CheckUserByID(id int) (result *gorm.DB, err error) {
	var user view.User
	result = db.MainDB.Where("ID = ?", id).First(&user)
	return
}
