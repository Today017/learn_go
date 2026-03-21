package main

import (
	"database/sql"
	"dbsample/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	/*
		got mod init dbsample
		got get -u github.com/go-sql-driver/mysql

		MySQLドライバ
		MySQLサーバーと直接通信する機能は別で持ってくる必要がある（database/sqlは抽象化されて実装されている）
	*/)

func main() {
	//データベースへの接続情報（ハードコーディングは良くないので後で直す）
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"

	//データベースに接続するためのアドレス分
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn) //データベースに接続 MySQLを使うよ〜
	//db: sql.DB型
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	{
		const sqlStr = `
			select title, contents, username, nice, created_at
			from articles
		`
		// fmtパッケージ：SQLインジェクション対策がされていないので良くない

		rows, err := db.Query(sqlStr)
		// rows: sql.Rows
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		articleArray := make([]models.Article, 0)
		for rows.Next() { // while と同じ感じ？
			var article models.Article
			// // 指定した変数ポインタに読み出す
			// err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

			var createdTime sql.NullTime
			err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

			if createdTime.Valid {
				article.CreatedAt = createdTime.Time
			}

			if err != nil {
				fmt.Println(err)
			} else {
				articleArray = append(articleArray, article)
			}
		}

		fmt.Printf("%+v\n", articleArray)
		fmt.Println("==================")
	}

	{ // 1行だけ読み出す場合の書き方
		articleID := 1000
		const sqlStr = `
			select *
			from articles
			where article_id = ?;
		`
		row := db.QueryRow(sqlStr, articleID)
		if err := row.Err(); err != nil {
			fmt.Println(err)
			return
		}

		var article models.Article
		var createdTime sql.NullTime
		err = row.Scan(&articleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
		if err != nil {
			fmt.Println(err)
			return
		}

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		fmt.Printf("%+v\n", article)
	}
}
