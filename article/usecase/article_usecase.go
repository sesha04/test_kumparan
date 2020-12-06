package usecase

import (
	"context"

	"github.com/sesha04/test_kumparan/domain"
)

type articleUsecase struct {
	articleRepo domain.ArticleRepository
}

func NewArticleUsecase(a domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo: a,
	}
}

func (au *articleUsecase) Create(ctx context.Context, a *domain.Article) (err error) {
	return au.articleRepo.Store(ctx, a)
}

func (au *articleUsecase) GetArticles(ctx context.Context, author, query string) ([]domain.Article, error) {
	q := domain.ArticleQuery{
		Author: author,
		Query:  query,
	}
	return au.articleRepo.GetArticles(ctx, q)
}
