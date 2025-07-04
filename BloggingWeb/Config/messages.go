package config

var Message = map[string]string{
	"ErrorWhileReadingPayloadData":  "Invalid request payload. Please provide correct data.",
	"ErrorWhileConnectingDB":        "Unable to connect to the database. Please try again later.",
	"ErrorWhileCheckUser":           "An error occurred while checking user existence in the database.",
	"ErrorWhileUploadingBlog":       "Failed to upload the blog post. Please try again.",
	"ErrorWhileGettingAllBlogs":     "Failed to get blogs.Please try again",
	"ErrorWhileConvertingStrToInt":  "Error While converting string value into int",
	"ErrorWhileGettingBlogByID":     "Please enter a valid blog id, unable to get blog by id",
	"ErrorWhileGettingBlogsByUid":   "Unable to get blogs by user id, Please try again",
	"ErrorWhileUpdatingBlogContent": "Error while updating blog content, Please try again",
	"ErrorWhileGettingAllUsers":     "Unable to get all users details, Please try again",
	"ErrorWhileGettingUserById":     "Error while getting user by id, Please try again",
	"PesatoAuthTokenErrMsg":         "Unable to get Pesato Auth Token, Please try again",
	"InvalidToken":                  "Pease enter a valid token",
}
