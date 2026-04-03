package repositories_test

import (
	"testing"

	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
	"github.com/Today017/learn_go/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				// Error: テスト失敗・処理は続行
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}

		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d article\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "syaku8",
	}

	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}

	if newArticle.ID <= 0 {
		t.Errorf("new article id is expected >=0 but got %d\n", newArticle.ID)
	}
	if newArticle.Title != article.Title {
		t.Errorf("article title is expected %q but got %q\n", article.Title, newArticle.Title)
	}
	if newArticle.Contents != article.Contents {
		t.Errorf("article contents is expected %q but got %q\n", article.Contents, newArticle.Contents)
	}
	if newArticle.UserName != article.UserName {
		t.Errorf("article user name is expected %q but got %q\n", article.UserName, newArticle.UserName)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ?;
		`
		testDB.Exec(sqlStr, article.Title)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	testID := 1

	err := repositories.UpdateNiceNum(testDB, testID)

	if err != nil {
		t.Error(err)
	}

	expectedNiceNum := 4
	article, err := repositories.SelectArticleDetail(testDB, testID)

	if article.NiceNum != expectedNiceNum {
		t.Errorf("article nice num is expected %d but got %d\n", expectedNiceNum, article.NiceNum)
	}

	t.Cleanup(func() {
		const sqlUpdateNice = `
			update articles
			set nice = ? where article_id = ?;
		`

		testDB.Exec(sqlUpdateNice, expectedNiceNum-1, testID)
	})
}
