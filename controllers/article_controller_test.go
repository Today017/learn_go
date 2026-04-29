package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

/*
GET /article/list 用のメソッドのテスト
クエリパラメータに数値が来た場合→200 OK のレスポンスコードが得られる
クエリパラメータに数値でないものが来た場合→400 BadRequest のレスポンスコードが得られる
*/
func TestArtcleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		reslutCode int
	}{
		{name: "number query", query: "1", reslutCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", reslutCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハンドラに渡す2つの引数を準備
			// w http.ResponseWriter
			// req *http.Request
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			// Post, Put のようにリクエストボディが必要ではないので nil でOK。
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			aCon.ArticleListHandler(res, req)

			if res.Code != tt.reslutCode {
				t.Errorf("unexpected StatusCOde: want %d but %d\n", tt.reslutCode, res.Code)
			}
		})
	}
}

/*
GET /article/{:id} 用のメソッドのテスト
パスパラメータ id に数値が来た場合→200 OK のレスポンスコードが得られる
パスパラメータ id に数値でないものが来た場合→400 BadRequest のレスポンスコードが得られる
*/
func TestArticleDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		articleID  string
		reslutCode int
	}{
		{name: "number pathparam", articleID: "1", reslutCode: http.StatusOK},
		{name: "alphabet pathparam", articleID: "aaa", reslutCode: http.StatusNotFound}, // gorilla/mux では正規表現に合わないURLには404が返ってくる
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()
			// ArticleDetailHandler では mux.Vars(req) でパスパラメータを取得するが、
			// gorilla/mux の仕様上、gorilla/mux のルータ経由で受け取ったリクエストでしか動作しない
			// よって、以下のように直接呼び出すコードだとうまくいかない
			// aCon.ArticleDetailHandler(res, req)

			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			// ルータ r 経由でリクエストを送信する
			r.ServeHTTP(res, req)

			if res.Code != tt.reslutCode {
				// t.Errorf("%s %s", res.Body, res.Result().Proto)
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.reslutCode, res.Code)
			}
		})
	}
}
