package main

import (
	"log"
	"net/http"

	"github.com/Today017/learn_go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//.Methods(...)でHTTPメソッドを制限
	//ハンドラ関数内でのメソッドチェックが不要になり、違うのが来たら自動で405を返してくれる
	//% curl http://localhost:8080/hello -x POST
	//curl: (5) Could not resolve proxy: POST
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

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
