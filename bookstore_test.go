package main

import (
	"github.com/danielgtaylor/huma/v2/humatest"
	c "go-bookstore/config"
	"testing"
)

func NewTestService(t *testing.T) humatest.TestAPI {
	_, api := humatest.New(t)
	db := c.NewSQLiteConnection(":memory:")
	c.AddRoutes(api, db)

	return api
}

func make_test_book() map[string]string {
	test_book := make(map[string]string)
	test_book["title"] = "the title"
	test_book["author"] = "the author"
	return test_book

}

func TestAddBook(t *testing.T) {

	api := NewTestService(t)
	test_book := make_test_book()

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
