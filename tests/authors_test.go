package tests

import (
	"encoding/json"
	m "go-bookstore/models"
	"testing"
)

func TestGetOneAuthor(t *testing.T) {

	api := NewTestService(t)
	author := map[string]any{"first_name": "john", "last_name": "doe"}

	resp := api.Post("/authors", author)

	var createdAuthor, retrievedAuthor *m.AuthorOutputRecord
	json.NewDecoder(resp.Body).Decode(&createdAuthor)

	resp = api.Get("/authors/" + createdAuthor.Id.String())
	json.NewDecoder(resp.Body).Decode(&retrievedAuthor)

	if retrievedAuthor.Id != createdAuthor.Id {
		t.Fatalf("Unexpected retrieved author. Created %v, retrieved %v", createdAuthor, retrievedAuthor)

	}

}

func TestGetAllAuthors(t *testing.T) {

	api := NewTestService(t)
	first := map[string]any{"first_name": "john", "last_name": "doe"}
	second := map[string]any{"first_name": "willy", "last_name": "wonka"}

	api.Post("/authors", first)
	api.Post("/authors", second)
	resp := api.Get("/authors")

	if resp.Code != 200 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}

}
