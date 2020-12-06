package domain

import (
	"context"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleQuery represent the query to get articles
type ArticleQuery struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Create(context.Context, *Article) error
	GetArticles(ctx context.Context, author, query string) ([]Article, error)
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Store(ctx context.Context, a *Article) error
	GetArticles(ctx context.Context, q ArticleQuery) ([]Article, error)
}
