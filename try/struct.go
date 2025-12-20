package main

import (
	"fmt"
	"time"
)

type Comment struct {
	commentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID          int
	Title       string
	Contents    string
	UserName    string
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

func main() {
	comment1 := Comment{
		commentID: 1,
		ArticleID: 1,
		Message:   "test comment",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		commentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "this is the first article",
		UserName:    "soma",
		NiceNum:     0,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	fmt.Printf("%+v\n", article)
}
