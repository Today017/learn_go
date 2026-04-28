package controllers_test

import (
	"testing"

	"github.com/Today017/learn_go/controllers"
	"github.com/Today017/learn_go/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

// テストで使うグローバル変数の準備
func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
