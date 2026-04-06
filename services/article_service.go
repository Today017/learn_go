package services

import (
	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
)

// GetArticleService
// 記事のIDを受け取る
// 該当する記事とエラーを返す
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticeService
// 記事を受け取ってデータベースに挿入し、挿入されたデータを返す
func (s *MyAppService) PostArticeService(article models.Article) (models.Article, error) {
	insertedArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return insertedArticle, nil
}

// GetArticleListService
// ページ番号を受け取って、そのページの記事一覧を返す
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

// PostNiceService
// 記事を受け取って、その記事のいいね数を+1して、更新された記事を返す
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	article.NiceNum++
	return article, nil
}
