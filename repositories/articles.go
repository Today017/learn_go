package repositories

import (
	"database/sql"

	"github.com/Today017/learn_go/models"
)

// 新規投稿をデータベースにinsertする
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles
		(title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
	`

	result, err := db.Exec(
		sqlStr,
		article.Title, article.Contents, article.UserName,
	)
	if err != nil {
		return models.Article{}, err
	}

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName =
		article.Title, article.Contents, article.UserName

	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

// 変数pageで指定されたページに表示する投稿一覧をデータベースから取得する
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		fromt articles
		limit ? offset ?;
	`

	artcileNumPerPage := 5
	offset := (artcileNumPerPage - 1) * page

	rows, err := db.Query(sqlStr, artcileNumPerPage, offset)
	if err != nil {
		return []models.Article{}, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime

		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Contents,
			&article.UserName,
			createdTime,
		)
		if err != nil {
			continue
		}

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// 投稿IDを指定して記事データを取得する
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(
		&articleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime,
	)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// いいねの数をupdateする関数
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `
		update articles
		set nice = ? where article_id = ?;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
