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

func (db Database) AddUserDetails(user *view.User) (err error) {
	result := db.MainDB.Table("users").Where("email = ?", user.Email).First(user)
	if result.RowsAffected == 0 {
		return db.MainDB.Table("users").Create(user).Error
	}
	return result.Error
}

func (db Database) CheckUserByID(blog view.Blog) (uid int, err error) {
	var user view.User
	result := db.MainDB.Table("users").Where("ID = ?", blog.UserID).First(&user)
	if result.RowsAffected == 0 {
		newUser := view.User{
			ID:            blog.UserID,
			FullName:      blog.FullName,
			FirstName:     blog.FirstName,
			LastName:      blog.LastName,
			Email:         blog.Email,
			Role:          "writer",
			BlogsUploaded: 1,
		}
		err = db.MainDB.Table("users").Create(&newUser).Error
		if err != nil {
			return user.ID, err
		}
		return newUser.ID, nil
	}
	userDlts := map[string]interface{}{
		"blogs_uploaded": user.BlogsUploaded + 1,
	}
	fmt.Println("users :", user)
	err = db.UpdateUserDetails(userDlts, user.ID)
	return user.ID, err
}

func (db Database) GetUserCredByEmail(email string) (user view.UserCredentials, err error) {
	err = db.MainDB.Where("email = ?", email).First(&user).Error
	return
}

func (db Database) SaveUserCredentials(user *view.UserCredentials) {
	db.MainDB.Table("user_credentials").Create(&user)

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
			"status":     true,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no session found for email %s", userCred.Email)
	}
	return
}

func (db Database) UpdateUserStatus(status bool, uid int) (err error) {
	fmt.Println("User id :", uid)
	result := db.MainDB.Table("users").
		Where("ID = ?", uid).
		Update("status", status)
	return result.Error
}
func (db Database) UpdateUserDetails(userDlts map[string]interface{}, uid int) (err error) {
	result := db.MainDB.Table("users").Where("ID= ?", uid).Updates(userDlts)
	return result.Error
}

func (db Database) UpdateUserSessionStatus(uid int) error {
	result := db.MainDB.Table("user_sessions").Where("uid= ?", uid).Update("status", false)
	return result.Error
}
