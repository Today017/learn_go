package repositories_test

import (
	"testing"

	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	tests := []struct {
		testTitle string
		articleID int
		expected  []models.Comment
	}{
		{
			testTitle: "subtest1",
			articleID: models.Article1.ID,
			expected: []models.Comment{
				{
					CommentID: 1,
					ArticleID: 1,
					Message:   "1st commnet yeah",
				}, {
					CommentID: 2,
					ArticleID: 1,
					Message:   "welcome",
				},
			},
		}, {
			testTitle: "subtest2",
			articleID: models.Article2.ID,
			expected:  []models.Comment{},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectCommentList(testDB, test.articleID)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) != len(test.expected) {
				t.Errorf("length of list is expected %d but got %d\n", len(test.expected), len(got))
			}

			judge := func(lhs models.Comment, rhs models.Comment) bool {
				if lhs.CommentID != rhs.CommentID {
					t.Errorf("comment id is expected %d but got %d\n", rhs.CommentID, lhs.CommentID)
					return false
				}
				if lhs.ArticleID != rhs.ArticleID {
					t.Errorf("article id is expected %d but got %d\n", rhs.ArticleID, lhs.ArticleID)
					return false
				}
				if lhs.Message != rhs.Message {
					t.Errorf("message is expected %q but got %q\n", rhs.Message, lhs.Message)
					return false
				}
				// if lhs.CreatedAt != rhs.CreatedAt {
				// 	t.Errorf("created at is expected %v but got %v\n", rhs.CreatedAt, lhs.CreatedAt)
				// 	return false
				// }
				return true
			}

			for i := 0; i < len(got); i++ {
				if !judge(got[i], test.expected[i]) {
					t.Errorf("not matched\n")
				}
			}
		})
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 2,
		Message:   "testcommentcomment",
	}

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID <= 0 {
		t.Errorf("new comment id is expected >=0 but got %d\n", newComment.CommentID)
	}
	if newComment.Message != comment.Message {
		t.Errorf("message is expected %q but got %q\n", comment.Message, newComment.Message)
	}
	if newComment.ArticleID != comment.ArticleID {
		t.Errorf("article id is expected %d but got %d\n", comment.ArticleID, newComment.ArticleID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?;
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}
