package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
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

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
