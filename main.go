package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Today017/learn_go/controllers"
	"github.com/Today017/learn_go/services"
	"github.com/gorilla/mux"

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
	con := controllers.NewMyAppController(ser)
	r := mux.NewRouter()

	//.Methods(...)でHTTPメソッドを制限
	//ハンドラ関数内でのメソッドチェックが不要になり、違うのが来たら自動で405を返してくれる
	//% curl http://localhost:8080/hello -x POST
	//curl: (5) Could not resolve proxy: POST

	// r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

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
