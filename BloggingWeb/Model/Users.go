package models

import (
	"fmt"

	"gorm.io/gorm"

	view "BloggingWeb/View"

)

func (db Database) GetAllUsers(user *[]view.User) error {
	return db.MainDB.Find(user).Error
}

func (db Database) GetUserByUid(user interface{}, uid int) (err error) {
	err = db.MainDB.First(user, uid).Error
	return
}
func (db Database) GetUserByEmail(user *view.User, email string) error {
    return db.MainDB.Where("email = ?", email).First(user).Error
}

func (db Database) AddUserDetails(user *view.User) {
	db.MainDB.Create(user)

}

func (db Database) CheckUserByID(id int) (result *gorm.DB, err error) {
	var user view.User
	result = db.MainDB.Where("ID = ?", id).First(&user)
	return
}

func (db Database) GetUserCredByEmail(email string) (user view.UserCredentials, err error) {
	err = db.MainDB.Where("email = ?", email).First(&user).Error
	return
}

func (db Database) SaveUserCredentials(user *view.UserCredentials) {
	db.MainDB.Create(&user)

}

func (db *Database) SaveSession(session *view.UserSession) error {
	return db.MainDB.Create(&session).Error
}

func (db *Database) UpdateSessionTokenAndAgent(userCred view.UserSession) (err error) {
	result := db.MainDB.Model(&view.UserSession{}).
		Where("email = ?", userCred.Email).
		Updates(map[string]interface{}{
			"token":      userCred.Token,
			"user_agent": userCred.UserAgent,
			"created_at": gorm.Expr("NOW()"),
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no session found for email %s", userCred.Email)
	}
	return
}
