package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/services"
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
	// ここで、reqArticleにJSONのなかみをデコードする
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil { // &はなんだっけ
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostArticeService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
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

	articles, err := services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// 記事ごと引数で受け取る形でいいの？
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail intenal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

//main.go->handlers/handler.go
//関数の定義の仕方変更
//先頭を大文字
//大文字にすることでパッケージ外からも参照可能にする
