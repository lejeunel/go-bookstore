package tests

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	goose "github.com/pressly/goose/v3"
	a "go-bookstore/app"
	sql "go-bookstore/repositories/sql"
	routes "go-bookstore/routes"
	s "go-bookstore/services"
	"testing"
)

func NewTestService(t *testing.T) humatest.TestAPI {
	_, api := humatest.New(t)
	db := a.NewSQLiteConnection(":memory:")

	goose.SetDialect(string(goose.DialectSQLite3))
	err := goose.Up(db.DB, "../migrations")
	if err != nil {
		panic(err)
	}
	bookRepo := sql.NewSQLBookRepo(db)
	authorRepo := sql.NewSQLAuthorRepo(db)

	bookService := s.BookService{BookRepo: bookRepo, AuthorRepo: authorRepo, MaxPageSize: 2,
		DefaultPageSize: 2}

	authorService := s.AuthorService{AuthorRepo: authorRepo, MaxPageSize: 2,
		DefaultPageSize: 2}
	routes.AddRoutes(api, "", bookService, authorService)

	return api
}
