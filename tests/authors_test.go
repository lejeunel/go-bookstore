package tests

import (
	"encoding/json"
	"fmt"
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

func TestGetAuthorPaginated(t *testing.T) {
	api := NewTestService(t)
	nAuthors := 20

	for i := 0; i < nAuthors; i++ {
		first_name := fmt.Sprintf("first_name %d", i)
		last_name := fmt.Sprintf("last_name %d", i)
		author := map[string]any{"first_name": first_name,
			"last_name": last_name}
		api.Post("/authors", author)
	}

	var retrievedNAuthors int
	var nextPage int = 1
	for {
		var results *m.AuthorPaginatedOutputBody
		url := fmt.Sprintf("/authors?page=%d", nextPage)
		resp := api.Get(url)
		json.NewDecoder(resp.Body).Decode(&results)
		retrievedNAuthors += len(results.Data)
		nextPage = results.Pagination.Next

		if nextPage == 0 {
			break
		}
	}
	if retrievedNAuthors != nAuthors {
		t.Fatalf("Unexpected retrieved num of authors. Created %v, retrieved %v", nAuthors, retrievedNAuthors)

	}

}
