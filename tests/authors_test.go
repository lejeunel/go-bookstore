package tests

import (
	"fmt"
	g "go-bookstore/generic"
	m "go-bookstore/models"
	"testing"
)

func TestAuthorPagination(t *testing.T) {
	s, ctx := NewTestServices(t)

	nAuthors := 20

	for i := 0; i < nAuthors; i++ {
		author := m.Author{FirstName: fmt.Sprintf("first_name %d", i),
			LastName: fmt.Sprintf("last_name %d", i)}
		s.Authors.Create(ctx, &author)
	}

	var retrievedNAuthors int
	var nextPage int = 1
	for {
		authors, paginationMeta, _ := s.Authors.GetOnePage(ctx, g.PaginationParams{Page: nextPage, PageSize: 2})
		retrievedNAuthors += len(authors)
		nextPage = paginationMeta.Next

		if paginationMeta.Next == 0 {
			break
		}
	}
	if retrievedNAuthors != nAuthors {
		t.Fatalf("Unexpected retrieved num of authors. Created %v, retrieved %v", nAuthors, retrievedNAuthors)

	}

}

func TestRetrievingDeletedAuthorShouldFail(t *testing.T) {
	s, ctx := NewTestServices(t)

	author := m.Author{FirstName: "a", LastName: "b"}
	s.Authors.Create(ctx, &author)
	s.Authors.Delete(ctx, author.Id.String())
	_, err := s.Authors.GetOne(ctx, author.Id.String())

	AssertError(t, err)
}

func TestDateOfBirthValidation(t *testing.T) {
	s, ctx := NewTestServices(t)

	author := m.Author{FirstName: "a", LastName: "b", DateOfBirth: "a long time ago"}
	_, err := s.Authors.Create(ctx, &author)

	AssertError(t, err)
}

func TestDeleteAssignedAuthorShouldFail(t *testing.T) {
	s, ctx := NewTestServices(t)

	book := m.Book{Title: "a"}
	author := m.Author{FirstName: "a", LastName: "b"}
	s.Authors.Create(ctx, &author)
	s.Books.Create(ctx, &book)
	s.Books.AssignAuthorToBook(ctx, book.Id.String(), author.Id.String())
	err := s.Authors.Delete(ctx, author.Id.String())

	AssertError(t, err)
}

func TestDeleteAuthor(t *testing.T) {
	s, ctx := NewTestServices(t)

	author := m.Author{FirstName: "a", LastName: "b"}
	s.Authors.Create(ctx, &author)
	err := s.Authors.Delete(ctx, author.Id.String())
	AssertNoError(t, err)

}
