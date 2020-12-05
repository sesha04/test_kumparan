package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sesha04/test_kumparan/article/usecase"
	"github.com/sesha04/test_kumparan/domain"
	"github.com/sesha04/test_kumparan/domain/mocks"
)

func TestCreate(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := domain.Article{
		Title:  "Hello",
		Body:   "Content",
		Author: "Sesha Andipa",
	}

	t.Run("success", func(t *testing.T) {
		tempMockArticle := mockArticle
		tempMockArticle.ID = 0
		mockArticleRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Article")).Return(nil).Once()

		u := usecase.NewArticleUsecase(mockArticleRepo)

		err := u.Create(context.TODO(), &tempMockArticle)

		assert.NoError(t, err)
		assert.Equal(t, mockArticle.Title, tempMockArticle.Title)
		mockArticleRepo.AssertExpectations(t)
	})

}
