package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

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

	now := time.Now()
	res, err := stmt.ExecContext(ctx, a.Title, a.Body, a.Author, now, now)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	a.CreatedAt = now
	a.UpdatedAt = now
	return
}

func (ar *articleRepository) GetArticles(ctx context.Context, q domain.ArticleQuery) ([]domain.Article, error) {
	queryList := []string{}
	args := []interface{}{}
	if q.Author != "" {
		queryList = append(queryList, "author = ?")
		args = append(args, q.Author)
	}
	if q.Query != "" {
		queryList = append(queryList, "(body LIKE ? OR title = ?)")
		args = append(args, fmt.Sprint("%", q.Query, "%"), q.Query)
	}
	where := ""
	if len(queryList) > 0 {
		where = "WHERE " + strings.Join(queryList, " AND ")
	}
	order := " ORDER BY created_at DESC"
	s := `SELECT id, title, body, author, updated_at, created_at FROM article`
	query := strings.Join([]string{s, where, order}, " ")
	rows, err := ar.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Article, 0)
	for rows.Next() {

		t := domain.Article{}
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Body,
			&t.Author,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, err
}
