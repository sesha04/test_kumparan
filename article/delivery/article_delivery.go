package delivery

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/sesha04/test_kumparan/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
	Cache    *cache.Cache
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase, ac *cache.Cache) {
	handler := &ArticleHandler{
		AUsecase: us,
		Cache:    ac,
	}
	e.POST("/articles", handler.Store)
	e.GET("/articles", handler.GetArticles)
}

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *ArticleHandler) Store(c echo.Context) (err error) {
	var article domain.Article
	err = c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.AUsecase.Create(ctx, &article)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, article)
}

func (a *ArticleHandler) GetArticles(c echo.Context) (err error) {
	ctx := c.Request().Context()
	author := c.Request().URL.Query().Get("author")
	query := c.Request().URL.Query().Get("query")
	ck := fmt.Sprint(author, "|", query)
	cp, exists := a.Cache.Get(ck)
	if exists {
		return c.JSON(http.StatusOK, cp)
	}
	result, err := a.AUsecase.GetArticles(ctx, author, query)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	defer func(result []domain.Article) {
		a.Cache.Set(ck, result, cache.DefaultExpiration)
	}(result)

	return c.JSON(http.StatusOK, result)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
