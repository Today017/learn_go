package repositories_test

import (
	"testing"

	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		exptected models.Article
	}{
		{
			testTitle: "subtest1",
			exptected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "soma",
				NiceNum:  3,
			},
		}, {
			testTitle: "subtest2",
			exptected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "soma",
				NiceNum:  4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.exptected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.exptected.ID {
				// Error: テスト失敗・処理は続行
				t.Errorf("ID: get %d but want %d\n", got.ID, test.exptected.ID)
			}
			if got.Title != test.exptected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.exptected.Title)
			}
			if got.Contents != test.exptected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.exptected.Contents)
			}
			if got.UserName != test.exptected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.exptected.UserName)
			}
			if got.NiceNum != test.exptected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.exptected.NiceNum)
			}

		})
	}
}

func TestSelectArticleList(t *testing.T) {
	exptectedNum := 5
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != exptectedNum {
		t.Errorf("want %d but got %d article\n", exptectedNum, num)
	}
}
