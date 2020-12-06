package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sesha04/test_kumparan/article/repository"
	"github.com/sesha04/test_kumparan/domain"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	ar := &domain.Article{
		Title:  "Judul",
		Body:   "Body",
		Author: "Sesha Andipa",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT  article SET title=\\? , body=\\? , author=\\?, updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Body, ar.Author, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(12, 1))

	a := repository.NewArticleRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}

func TestGetArticleWithoutParams(t *testing.T) {
	q := domain.ArticleQuery{}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Article{
		{
			ID: 1, Title: "title 1", Body: "Body 1",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		{
			ID: 2, Title: "title 2", Body: "Body 2",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "body", "author", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Body,
			mockArticles[0].Author, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Body,
			mockArticles[1].Author, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	query := "SELECT id, title, body, author, updated_at, created_at FROM article ORDER BY created_at DESC"
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := repository.NewArticleRepository(db)
	list, err := a.GetArticles(context.TODO(), q)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetArticleByAuthor(t *testing.T) {
	author := "Sesha Andipa"
	q := domain.ArticleQuery{
		Author: author,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Article{
		{
			ID: 1, Title: "title 1", Body: "Body 1",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		{
			ID: 2, Title: "title 2", Body: "Body 2",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "body", "author", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Body,
			mockArticles[0].Author, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Body,
			mockArticles[1].Author, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	query := "SELECT id, title, body, author, updated_at, created_at FROM article WHERE author = \\? ORDER BY created_at DESC"
	mock.ExpectQuery(query).WithArgs(q.Author).WillReturnRows(rows)

	a := repository.NewArticleRepository(db)
	list, err := a.GetArticles(context.TODO(), q)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetArticleByBodyOrTitle(t *testing.T) {
	words := "word1 word2"
	q := domain.ArticleQuery{
		Query: words,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Article{
		{
			ID: 1, Title: "title 1", Body: "word1 word2",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		{
			ID: 2, Title: "title 2", Body: "word1 word2 word3",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "body", "author", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Body,
			mockArticles[0].Author, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Body,
			mockArticles[1].Author, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	query := "SELECT id, title, body, author, updated_at, created_at FROM article WHERE \\(body LIKE \\? OR title = \\?\\) ORDER BY created_at DESC"
	mock.ExpectQuery(query).WithArgs(fmt.Sprint("%", q.Query, "%"), q.Query).WillReturnRows(rows)

	a := repository.NewArticleRepository(db)
	list, err := a.GetArticles(context.TODO(), q)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetArticleByBodyOrTitleAndAuthor(t *testing.T) {
	author := "Sesha Andipa"
	words := "word1 word2"
	q := domain.ArticleQuery{
		Author: author,
		Query:  words,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Article{
		{
			ID: 1, Title: "title 1", Body: "word1 word2",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		{
			ID: 2, Title: "title 2", Body: "word1 word2 word3",
			Author: "Sesha Andipa", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "body", "author", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Body,
			mockArticles[0].Author, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Body,
			mockArticles[1].Author, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	query := "SELECT id, title, body, author, updated_at, created_at FROM article WHERE author = \\? AND \\(body LIKE \\? OR title = \\?\\) ORDER BY created_at DESC"
	mock.ExpectQuery(query).WithArgs(q.Author, fmt.Sprint("%", q.Query, "%"), q.Query).WillReturnRows(rows)

	a := repository.NewArticleRepository(db)
	list, err := a.GetArticles(context.TODO(), q)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
