package main

import (
	"database/sql"
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

	//Pingメソッド: データベースへのコネクションが生きているかの確認
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}
