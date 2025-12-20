package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Today017/learn_go/models"
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
	article := models.Article1

	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
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

	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		errMsg := fmt.Sprintf("Fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)

		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("Fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1

	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1

	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "Fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

//main.go->handlers/handler.go
//関数の定義の仕方変更
//先頭を大文字
//大文字にすることでパッケージ外からも参照可能にする
