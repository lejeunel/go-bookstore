package main

import (
	"encoding/json"
	"fmt"
	"github.com/danielgtaylor/huma/v2/humatest"
	"go-bookstore/config/driver"
	"strings"
	"testing"
)

func NewTestService(t *testing.T) humatest.TestAPI {
	_, api := humatest.New(t)
	db := driver.NewSQLiteConnection(":memory:")
	addRoutes(api, db)

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

	// Convert the JSON response to a map
	var response map[string]string
	// Grab the value & whether or not it exists
	json.Unmarshal([]byte(resp.Body.String()), &response)
	id, id_exists := response["id"]

	fmt.Println(response)

	if !id_exists {
		t.Fatalf("Expected to retrieve JSON with an id field, but got none")
	}
	resp2 := api.Get(fmt.Sprintf("/books/%s", id))
	if !strings.Contains(resp2.Body.String(), test_book["title"]) {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}

}

func TestGetBookWithWrongIdReturns404(t *testing.T) {
	api := NewTestService(t)

	resp := api.Get("/books/xyz")
	if resp.Code != 404 {
		t.Fatalf("Unexpected status code: %d, expected 404", resp.Code)
	}
}
