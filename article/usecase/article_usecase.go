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

func (au *articleUsecase) Create(ctx context.Context, a *domain.Article) (err error){
	err = au.articleRepo.Store(ctx, a)
	return err
}
