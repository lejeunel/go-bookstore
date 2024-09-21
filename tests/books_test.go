package tests

import (
	"encoding/json"
	"fmt"
	"github.com/danielgtaylor/huma/v2/humatest"
	goose "github.com/pressly/goose/v3"
	a "go-bookstore/app"
	m "go-bookstore/models"
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

func TestAddBook(t *testing.T) {

	api := NewTestService(t)

	var created_book *m.BookOutputRecord
	resp := api.Post("/books", map[string]any{"title": "the title"})
	json.NewDecoder(resp.Body).Decode(&created_book)

	if resp.Code != 201 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}

	var first_author, second_author *m.AuthorOutputRecord
	resp = api.Post("/authors", map[string]any{
		"first_name":    "john",
		"last_name":     "doe",
		"date_of_birth": "",
	})
	json.NewDecoder(resp.Body).Decode(&first_author)

	resp = api.Post("/authors", map[string]any{
		"first_name":    "jane",
		"last_name":     "smith",
		"date_of_birth": "",
	})
	json.NewDecoder(resp.Body).Decode(&second_author)

	resp = api.Post("/books/" + created_book.Id.String() + "/authors/" + first_author.Id.String())
	resp = api.Post("/books/" + created_book.Id.String() + "/authors/" + second_author.Id.String())

	var final_book *m.BookOutputRecord
	json.NewDecoder(resp.Body).Decode(&final_book)

	if len(final_book.Authors) != 2 {
		t.Fatalf("Expected to retrieve 2 associated authors, got %v", final_book.Authors)

	}

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

func TestGetBookWithWrongIdReturns404(t *testing.T) {
	api := NewTestService(t)

	resp := api.Get("/books/xyz")
	if resp.Code != 404 {
		t.Fatalf("Unexpected status code: %d, expected 404", resp.Code)
	}
}

func TestGetBookPaginated(t *testing.T) {
	api := NewTestService(t)
	nBooks := 20

	for i := 0; i < nBooks; i++ {
		title := fmt.Sprintf("book %d", i)
		book := map[string]any{"title": title}
		api.Post("/books", book)
	}

	var retrievedNBooks int
	var nextPage int = 1
	for {
		var results *m.BookPaginatedOutputBody
		url := fmt.Sprintf("/books?page=%d", nextPage)
		resp := api.Get(url)
		json.NewDecoder(resp.Body).Decode(&results)
		retrievedNBooks += len(results.Data)
		nextPage = results.Pagination.Next

		if nextPage == 0 {
			break
		}
	}
	if retrievedNBooks != nBooks {
		t.Fatalf("Unexpected retrieved num of books. Created %v, retrieved %v", nBooks, retrievedNBooks)

	}

}
