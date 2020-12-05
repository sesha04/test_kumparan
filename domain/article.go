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

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Create(context.Context, *Article) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Store(ctx context.Context, a *Article) error
}
