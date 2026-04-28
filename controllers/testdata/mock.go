package testdata

/*
	- PostArticleService(article models.Article) (models.Article, error)
	- GetArticleListService(page int) ([]models.Article, error)
	- GetArticleService(articleID int) (models.Article, error)
	- PostNiceService(article models.Article) (models.Article, error)
	- PostCommentService(comment models.Comment) (models.Comment, error)
*/

import "github.com/Today017/learn_go/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commnetTestData[0], nil
}
