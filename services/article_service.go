package services

import (
	"database/sql"
	"errors"

	"github.com/Today017/learn_go/apperrors"
	"github.com/Today017/learn_go/models"
	"github.com/Today017/learn_go/repositories"
)

// GetArticleService
// 記事のIDを受け取る
// 該当する記事とエラーを返す
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	// エラーハンドリング
	// 1. select 文の実行自体には成功したが、結果が 0 件のパターン
	// 2. select 文の実行自体に失敗したパターン

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
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
		err = apperrors.InsertDataFaild.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return insertedArticle, nil
}

// GetArticleListService
// ページ番号を受け取って、そのページの記事一覧を返す
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	// db.Query は検索結果が 0 件でも error が返ってこないので個別に判定
	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// PostNiceService
// 記事を受け取って、その記事のいいね数を+1して、更新された記事を返す
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdataDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	article.NiceNum++
	return article, nil
}
