package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	view "BloggingWeb/View"
)

func CreateBlog(c *gin.Context) {
	var blog view.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1001" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	// In this function we used to make connection with database.
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1002" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	// This function is used to disconnect the db connection.
	defer config.DisconnectDbConnection(database.MainDB)
	result, err := database.CheckUserByID(blog.UserID)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1003" + config.Message["ErrorWhileCheckUser"], Error: err})
	}
	// In this step check, number of rows that were matched by the database. And if user is not present then add user into db.
	if result.RowsAffected == 0 {
		user := view.User{FullName: blog.FullName, FirstName: blog.FirstName, LastName: blog.LastName, Email: blog.Email, ID: blog.UserID, Role: "writer"}
		database.AddUserDetails(&user)
		blog.UserID = user.ID
	}
	if err := database.CreatePost(&blog); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1004" + config.Message["ErrorWhileUploadingBlog"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: config.Message["BlogUploadedSuccessfully"], Response: blog})
}

func GetBlogs(c *gin.Context) {
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1005" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	var posts []view.Blog
	if err := database.GetAllPosts(&posts); err != nil {
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
	defer config.DisconnectDbConnection(database.MainDB)
	var blog view.Blog
	if err := database.GetBlogByID(&blog, id); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1009" + config.Message["ErrorWhileGettingBlogByID"], Error: err})
		return
	}
	database.DeleteTable(&blog)
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
	defer config.DisconnectDbConnection(database.MainDB)
	var blog view.Blog
	if err := database.GetBlogByID(&blog, id); err != nil {
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
	defer config.DisconnectDbConnection(db.MainDB)
	var blogs []view.Blog
	if err := db.GetBlogsByUid(&blogs, uid); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1015" + config.Message["ErrorWhileGettingBlogsByUid"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: blogs})
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
	defer config.DisconnectDbConnection(db.MainDB)
	if err := db.UpdatingBlogData(&blog); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1017" + config.Message["ErrorWhileUpdatingBlogContent"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Successfully updated", Response: blog})
}
