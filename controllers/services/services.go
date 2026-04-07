package services

import "github.com/Today017/learn_go/models"

type MyAppServicer interface {
	PostArticeService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(models.Article, error)
	PostCommentService(comment models.Comment) (models.Comment, error)
}
