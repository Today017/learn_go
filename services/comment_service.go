package services

import (
	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
)

// PostCommentService
// コメントを受け取り、そのコメントをデータベースに挿入し、挿入したコメントを返す
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	insertedComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return insertedComment, nil
}
