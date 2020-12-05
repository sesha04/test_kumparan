package repository

import (
	"context"
	"database/sql"

	"github.com/sesha04/test_kumparan/domain"
)

type articleRepository struct {
	Conn *sql.DB
}

// NewArticleRepository will create an object that represent the article.Repository interface
func NewArticleRepository(Conn *sql.DB) domain.ArticleRepository {
	return &articleRepository{Conn}
}

func (ar *articleRepository) Store(ctx context.Context, a *domain.Article) (err error) {
	query := `INSERT  article SET title=? , body=? , author=?, updated_at=? , created_at=?`
	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Body, a.Author, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}
