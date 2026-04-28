package testdata

import "github.com/Today017/learn_go/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "soma",
		NiceNum:     2,
		CommentList: commnetTestData,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "soma",
		NiceNum:  4,
	},
}

var commnetTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st commnet yeah",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
