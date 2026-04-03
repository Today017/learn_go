package repositories

import (
	"database/sql"

	"github.com/Today017/learn_go/models"
)

// 新規投稿をデータベースにinsert
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	result, err := db.Exec(
		sqlStr,
		comment.ArticleID, comment.Message,
	)
	if err != nil {
		return models.Comment{}, err
	}

	var newComment models.Comment
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return []models.Comment{}, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		err := rows.Scan(
			&comment.CommentID,
			&comment.ArticleID,
			&comment.Message,
			&createdTime,
		)
		if err != nil {
			return []models.Comment{}, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
