package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}
	Article1 = Article{
		ID:          1,
		Title:       "soma🙂",
		Contents:    "soma soma soma",
		UserName:    "soma",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		ID:          2,
		Title:       "Syaku8😐",
		Contents:    "this is the second article",
		UserName:    "soma",
		NiceNum:     2,
		CommentList: []Comment{},
		CreatedAt:   time.Now(),
	}
)
