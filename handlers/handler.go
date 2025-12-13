package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, s)
	//io.Writer型wに文字列sを書き込む

	if req.Method == http.MethodGet { //ハードコーディングは避けよう
		io.WriteString(w, "Hello, world!\n")
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
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
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
	}
}

//main.go->handlers/handler.go
//関数の定義の仕方変更
//先頭を大文字
//大文字にすることでパッケージ外からも参照可能にする
