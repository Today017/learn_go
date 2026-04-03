package testdata

import "github.com/Today017/learn_go/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "soma",
		NiceNum:  3,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "soma",
		NiceNum:  4,
	},
}

var CommentTestData = []models.Comment{
	{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	}, {
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
