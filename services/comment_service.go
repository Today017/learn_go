package services

import (
	"github.com/Today017/learn_go/apperrors"
	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
)

// PostCommentService
// コメントを受け取り、そのコメントをデータベースに挿入し、挿入したコメントを返す
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	insertedComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFaild.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return insertedComment, nil
}
