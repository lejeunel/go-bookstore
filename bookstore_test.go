package main

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	c "go-bookstore/config"
	r "go-bookstore/repositories"
	"testing"
)

func NewTestService(t *testing.T) humatest.TestAPI {
	_, api := humatest.New(t)
	db := c.NewSQLiteConnection(":memory:")
	paginator := &r.Paginator{MaxPageSize: 2}
	c.AddRoutes(api, db, paginator, "")

	return api
}

func MakeTestBook() map[string]string {
	test_book := make(map[string]string)
	test_book["title"] = "the title"
	test_book["author"] = "the author"
	return test_book

}

func TestAddBook(t *testing.T) {

	api := NewTestService(t)
	test_book := MakeTestBook()

	resp := api.Post("/books", test_book)

	if resp.Code != 201 {
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
