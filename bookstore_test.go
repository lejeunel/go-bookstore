package main

import (
	"encoding/json"
	"fmt"
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

func TestGetBookPagination(t *testing.T) {
	api := NewTestService(t)

	test_book := make(map[string]string)
	test_book["author"] = "the author"

	for i := 1; i < 4; i++ {
		test_book["title"] = fmt.Sprintf("title_%03d", i)
		api.Post("/books", test_book)
	}

	res := api.Get("/books?page=1&pagesize=2")

	var m []map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if len(m) != 2 {
		t.Fatalf("Expected to retrieve 2 items. Got '%v'", len(m))
	}
}
