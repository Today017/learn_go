package handlers

import (
	"encoding/json"
	"io"
	"log"
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
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil { // &はなんだっけ
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
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

	log.Println(page)

	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := mux.Vars(req)["id"]
	log.Println(articleID)
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

//main.go->handlers/handler.go
//関数の定義の仕方変更
//先頭を大文字
//大文字にすることでパッケージ外からも参照可能にする
