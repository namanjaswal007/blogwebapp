package controller

import (
	"BloggingWeb/BlogDB"
	config "BloggingWeb/Config"
	view "BloggingWeb/View"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog view.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1001" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	// In this function we used to make connection with database.
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1002" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	// This function is used to disconnect the db connection.
	defer db.MainDB.Close()
	db.Query = BlogDB.New(db.MainDB)
	blog.UserID, err = db.SaveUserDtls(blog)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1003 " + config.Message["ErrorWhileCheckUser"], Error: err.Error()})
		return
	}
	postedBlog, err := db.Query.PostBlog(context.Background(), BlogDB.PostBlogParams{FirstName: blog.FirstName, LastName: blog.LastName, FullName: blog.FullName, Content: blog.Content, Title: blog.Title, UserID: int64(blog.UserID), Email: blog.Email})
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1004 " + config.Message["ErrorWhileUploadingBlog"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: config.Message["BlogUploadedSuccessfully"], Response: postedBlog})
}

func GetBlogs(c *gin.Context) {
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1005" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer database.MainDB.Close()
	query := BlogDB.New(database.MainDB)
	var posts []BlogDB.Blog
	if posts, err = query.GetAllBlogs(context.Background()); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1006" + config.Message["ErrorWhileGettingAllBlogs"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Success", Response: posts})
}

func DeleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1007" + config.Message["ErrorWhileConvertingStrToInt"], Error: err})
		return
	}
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1008" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer database.MainDB.Close()
	query := BlogDB.New(database.MainDB)
	if err := query.DeleteBlogByID(context.Background(), int64(id)); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1009" + config.Message["ErrorWhileGettingBlogByID"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Blog is deleted successfully"})
}

func GetBlogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1010" + config.Message["ErrorWhileConvertingStrToInt"], Error: err})
		return
	}
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1011" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer database.MainDB.Close()
	query := BlogDB.New(database.MainDB)
	var blog BlogDB.Blog
	if blog, err = query.GetBlogByID(context.Background(), int64(id)); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1012" + config.Message["ErrorWhileGettingBlogByID"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: blog})
}

func GetUserBlogs(c *gin.Context) {
	uid, err := strconv.Atoi(c.DefaultQuery("uid", ""))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1013" + config.Message["ErrorWhileConvertingStrToInt"], Error: err})
		return
	}
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1014" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer db.MainDB.Close()
	query := BlogDB.New(db.MainDB)
	var userBlogs []BlogDB.Blog
	if userBlogs, err = query.GetUserBlogs(context.Background(), int64(uid)); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1015" + config.Message["ErrorWhileGettingBlogsByUid"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: userBlogs})
}

func UpdateBlog(c *gin.Context) {
	var blog view.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1015" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1016" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer db.MainDB.Close()
	var updatedBlog BlogDB.Blog
	query := BlogDB.New(db.MainDB)
	if updatedBlog, err = query.UpdateBlogsContent(context.Background(), BlogDB.UpdateBlogsContentParams{Content: blog.Content, Title: blog.Title, BlogID: int64(blog.BlogId)}); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1017" + config.Message["ErrorWhileUpdatingBlogContent"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Successfully updated", Response: updatedBlog})
}
