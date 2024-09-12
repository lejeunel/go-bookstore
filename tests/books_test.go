package tests

import (
	"encoding/json"
	"github.com/danielgtaylor/huma/v2/humatest"
	goose "github.com/pressly/goose/v3"
	a "go-bookstore/app"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
	routes "go-bookstore/routes"
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
	paginator := &r.Paginator{MaxPageSize: 2}
	routes.AddRoutes(api, db, paginator, "")

	return api
}

func TestAddBook(t *testing.T) {

	api := NewTestService(t)
	test_book := map[string]any{"title": "the title"}

	resp := api.Post("/books", test_book)

	if resp.Code != 201 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}

	resp = api.Post("/authors", map[string]any{
		"first_name":    "john",
		"last_name":     "doe",
		"date_of_birth": "",
	})

}
func TestGetOneBook(t *testing.T) {

	api := NewTestService(t)
	book := map[string]any{"title": "the first title"}

	resp := api.Post("/books", book)

	var createdBook, retrievedBook *m.BookOutputRecord
	json.NewDecoder(resp.Body).Decode(&createdBook)

	resp = api.Get("/books/" + createdBook.Id.String())
	json.NewDecoder(resp.Body).Decode(&retrievedBook)

	if retrievedBook.Title != createdBook.Title {
		t.Fatalf("Unexpected retrieved book. Created %v, retrieved %v", createdBook, retrievedBook)

	}

}

func TestGetAllBooks(t *testing.T) {

	api := NewTestService(t)
	first := map[string]any{"title": "the first title"}
	second := map[string]any{"title": "the second title"}

	api.Post("/books", first)
	api.Post("/books", second)
	resp := api.Get("/books")

	if resp.Code != 200 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}

}

func TestGetBookWithWrongIdReturns404(t *testing.T) {
	api := NewTestService(t)

	resp := api.Get("/books/xyz")
	if resp.Code != 404 {
		t.Fatalf("Unexpected status code: %d, expected 404", resp.Code)
	}
}

func TestGetBookPaginationMinimum(t *testing.T) {
	api := NewTestService(t)

	resp := api.Get("/books?page=0&pagesize=2")

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d, expected 424", resp.Code)
	}
}
