package delivery_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	articleDelivery "github.com/sesha04/test_kumparan/article/delivery"
	"github.com/sesha04/test_kumparan/domain"
	"github.com/sesha04/test_kumparan/mocks"
)

func TestStore(t *testing.T) {
	mockArticle := domain.Article{
		Title:  "Title",
		Body:   "Body",
		Author: "Sesha Andipa",
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(mocks.ArticleUsecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Create", mock.Anything, mock.AnythingOfType("*domain.Article")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles")

	handler := articleDelivery.ArticleHandler{
		AUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStoreBadRequest(t *testing.T) {
	mockArticle := domain.Article{
		Title: "Title",
		Body:  "Body",
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(mocks.ArticleUsecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles")

	handler := articleDelivery.ArticleHandler{
		AUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	author := "Sesha Andipa"
	query := "sesuatu"
	mockUCase := new(mocks.ArticleUsecase)

	mockUCase.On("GetArticles", mock.Anything, author, query).Return([]domain.Article{}, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/articles?author=Sesha%20Andipa&query=sesuatu", nil)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles?author=Sesha%20Andipa&query=sesuatu")

	handler := articleDelivery.ArticleHandler{
		AUsecase: mockUCase,
	}
	err = handler.GetArticles(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
