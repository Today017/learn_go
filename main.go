package main

import (
	"log"
	"net/http"
	"github.com/Today017/learn_go/handlers"
)

func main() {
	//http.HandleFuncにハンドラ関数を渡すことで、サーバーでこのハンドラが使われるようになる
	http.HandleFunc("/hello", handlers.HelloHandler) //パス指定
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	//ターミナルへのログ表示
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//http.ListenAndServe(addr string, handler Handler) error
	//addr=アドレス
	//error=サーバー起動時に起こったエラー
	//errorをlog.Fatelに渡している
	//log.Fatal=エラーをログ出力して異常終了

	//handler=nilは？
}
