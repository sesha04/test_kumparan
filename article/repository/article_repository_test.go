package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sesha04/test_kumparan/article/repository"
	"github.com/sesha04/test_kumparan/domain"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &domain.Article{
		Title:     "Judul",
		Body:      "Body",
		CreatedAt: now,
		UpdatedAt: now,
		Author:    "Sesha Andipa",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT  article SET title=\\? , body=\\? , author=\\?, updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Body, ar.Author, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := repository.NewArticleRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}
