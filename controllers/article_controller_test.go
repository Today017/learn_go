package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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
