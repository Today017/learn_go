package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Today017/learn_go/controllers"
	"github.com/Today017/learn_go/routers"
	"github.com/Today017/learn_go/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	r := routers.NewRouter(aCon, cCon)

	//ターミナルへのログ表示
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	//http.ListenAndServe(addr string, handler Handler) error
	//addr=アドレス
	//error=サーバー起動時に起こったエラー
	//errorをlog.Fatalに渡している
	//log.Fatal=エラーをログ出力して異常終了

	//handler=nilは？ -> 伏線回収、か
	//サーバーの中で使うルーターを指定する部分 nilならGo標準のルーターが採用される
}
