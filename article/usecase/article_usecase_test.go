package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sesha04/test_kumparan/article/usecase"
	"github.com/sesha04/test_kumparan/domain"
	"github.com/sesha04/test_kumparan/mocks"
)

func TestCreate(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := domain.Article{
		Title:  "Hello",
		Body:   "Body",
		Author: "Sesha Andipa",
	}

	tempMockArticle := mockArticle
	mockArticleRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Article")).Return(nil).Once()
	u := usecase.NewArticleUsecase(mockArticleRepo)
	err := u.Create(context.TODO(), &tempMockArticle)

	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
}

func TestGetArticles(t *testing.T) {
	author := "Sesha Andipa"
	search := "sesuatu"
	query := domain.ArticleQuery{
		Author: author,
		Query:  search,
	}
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticleRepo.On("GetArticles", mock.Anything, query).Return([]domain.Article{}, nil).Once()
	u := usecase.NewArticleUsecase(mockArticleRepo)
	_, err := u.GetArticles(context.TODO(), author, search)

	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
}
