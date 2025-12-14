package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, s)
	//io.Writer型wに文字列sを書き込む

	io.WriteString(w, "Hello, world!\n")
	//io.Writer型=インターフェース
	/*
		type Write interface {
			Write(p []byte) (n int, err error)
		}
	*/
	//-> メソッドとしてWrite(int n, error err)->[]byte pを持つならOK
	//C++のコンセプトみたいなやつ？

}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil || page < 1 {
			http.Error(w, "Invalid query parameter (in handler)", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// articleID := 1
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)

		return
	}

	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}

//main.go->handlers/handler.go
//関数の定義の仕方変更
//先頭を大文字
//大文字にすることでパッケージ外からも参照可能にする
