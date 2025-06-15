package tests

import (
	ctx "context"
	"fmt"
	"github.com/danielgtaylor/huma/v2/humatest"
	goose "github.com/pressly/goose/v3"
	a "go-bookstore/app"
	sql "go-bookstore/repositories/sql"
	routes "go-bookstore/routes"
	s "go-bookstore/services"
	"testing"
)

type Services struct {
	Books   *s.BookService
	Authors *s.AuthorService
}

func NewTestServices(t *testing.T) (Services, ctx.Context) {
	db := a.NewSQLiteConnection(":memory:")
	goose.SetLogger(goose.NopLogger())
	goose.SetDialect(string(goose.DialectSQLite3))
	err := goose.Up(db.DB, "../migrations")
	if err != nil {
		panic(err)
	}
	bookRepo := sql.NewSQLBookRepo(db)
	authorRepo := sql.NewSQLAuthorRepo(db)

	bookService := s.NewBookService(bookRepo, authorRepo, 2, 2)
	authorService := s.NewAuthorService(authorRepo, bookRepo, 2, 2)

	return Services{Books: bookService, Authors: authorService}, ctx.Background()

}

func NewTestAPI(t *testing.T) humatest.TestAPI {
	_, api := humatest.New(t)

	testServices, _ := NewTestServices(t)
	routes.AddAPIRoutes(api, "", *testServices.Books, *testServices.Authors)

	return api
}

func AssertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error(fmt.Printf("did not want an error but got: %v", err))
	}
}
