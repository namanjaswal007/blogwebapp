package models

import (
	"BloggingWeb/BlogDB"
	view "BloggingWeb/View"
	"context"
	"database/sql"
)

type Database struct {
	MainDB *sql.DB
	Query  *BlogDB.Queries
}

func (db Database) AddUserDetails(user *view.User) (err error) {
	userDtls, err := db.Query.GetUserByEmail(context.Background(), user.Email)
	// var userDtls BlogDB.User
	if err != nil {
		userDtls, err = db.Query.InsertUser(context.Background(), BlogDB.InsertUserParams{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Role: sql.NullString{String: user.Role, Valid: true}, Status: sql.NullBool{Bool: true, Valid: true}})
		if err != nil {
			return
		}
	}
	user.UserId = int(userDtls.UserID)
	return
}

func (db Database) SaveUserDtls(blog view.Blog) (uid int, err error) {
	var user BlogDB.User
	exist, err := db.Query.CheckUserExists(context.Background(), blog.Email)
	if err != nil {
		return
	}
	var userId int64
	userId = int64(blog.UserID)
	if !exist {
		newUser := BlogDB.InsertUserParams{
			FullName:  blog.FullName,
			FirstName: blog.FirstName,
			LastName:  blog.LastName,
			Email:     blog.Email,
			Status:    sql.NullBool{Bool: false, Valid: true},
			Role:      sql.NullString{String: "writer", Valid: true},
		}
		user, err = db.Query.InsertUser(context.Background(), newUser)
		if err != nil {
			return int(user.UserID), err
		}
		userId = user.UserID
	}
	user, err = db.Query.UpdateUserBlogCount(context.Background(), BlogDB.UpdateUserBlogCountParams{BlogsUploaded: sql.NullInt32{Int32: 1, Valid: true}, UserID: userId})
	return int(user.UserID), err
}
