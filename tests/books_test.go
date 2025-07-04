package tests

import (
	"fmt"
	g "go-bookstore/generic"
	m "go-bookstore/models"
	"testing"
)

func TestAddBook(t *testing.T) {

	s, ctx := NewTestServices(t)

	book := m.Book{Title: "the title"}
	createdBook, _ := s.Books.Create(ctx, &book)
	retrievedBook, _ := s.Books.GetOne(ctx, createdBook.Id.String())

	if retrievedBook.Title != createdBook.Title {
		t.Fatalf("Expected to retrieve book with title %v got %v", createdBook.Title, retrievedBook.Title)
	}

}

func TestAssignAuthorsToBook(t *testing.T) {

	s, ctx := NewTestServices(t)

	book := m.Book{Title: "the title"}
	firstAuthor := m.Author{FirstName: "john", LastName: "doe", DateOfBirth: ""}
	secondAuthor := m.Author{FirstName: "jane", LastName: "smith", DateOfBirth: ""}

	s.Books.Create(ctx, &book)
	s.Authors.Create(ctx, &firstAuthor)
	s.Authors.Create(ctx, &secondAuthor)

	s.Books.AssignAuthorToBook(ctx, book.Id.String(), firstAuthor.Id.String())
	s.Books.AssignAuthorToBook(ctx, book.Id.String(), secondAuthor.Id.String())
	retrievedBook, _ := s.Books.GetOne(ctx, book.Id.String())

	if len(retrievedBook.Authors) != 2 {
		t.Fatalf("Expected to retrieve 2 associated authors, got %v", len(retrievedBook.Authors))
	}

	if retrievedBook.Authors[0].FirstName != firstAuthor.FirstName {
		t.Fatalf("Expected name of author %v, got %v", firstAuthor.FirstName, retrievedBook.Authors[0].FirstName)
	}

}

func TestDeleteBook(t *testing.T) {
	s, ctx := NewTestServices(t)

	book := m.Book{Title: "the title"}
	author := m.Author{FirstName: "john", LastName: "doe", DateOfBirth: ""}
	s.Books.Create(ctx, &book)
	s.Authors.Create(ctx, &author)

	s.Books.AssignAuthorToBook(ctx, book.Id.String(), author.Id.String())
	s.Books.Delete(ctx, book.Id.String())

	booksOfAuthor, err := s.Authors.GetBooksOfAuthor(ctx, &author)
	if err != nil {
		AssertNoError(t, err)
	}

	if len(booksOfAuthor) != 0 {
		t.Fatalf("Expected to retrieved 0 books, got %v", len(booksOfAuthor))
	}

}

func TestAssignAuthorToNonExistingBookShouldFail(t *testing.T) {
	s, ctx := NewTestServices(t)

	book := m.Book{Title: "the title"}
	author := m.Author{FirstName: "john", LastName: "doe", DateOfBirth: ""}
	s.Books.Create(ctx, &book)
	s.Authors.Create(ctx, &author)

	_, err := s.Books.AssignAuthorToBook(ctx, "bad-id", author.Id.String())

	AssertError(t, err)
}

func TestRetrievingDeletedBookShouldFail(t *testing.T) {
	s, ctx := NewTestServices(t)

	book := m.Book{Title: "the title"}
	s.Books.Create(ctx, &book)
	s.Books.Delete(ctx, book.Id.String())
	_, err := s.Books.GetOne(ctx, book.Id.String())

	AssertError(t, err)
}

func TestBookPagination(t *testing.T) {
	s, ctx := NewTestServices(t)

	nBooks := 20

	for i := 0; i < nBooks; i++ {
		title := fmt.Sprintf("book %d", i)
		book := m.Book{Title: title}
		s.Books.Create(ctx, &book)
	}

	var retrievedNBooks int
	var nextPage int = 1
	for {
		books, paginationMeta, _ := s.Books.GetOnePage(ctx, g.PaginationParams{Page: nextPage, PageSize: 2})
		retrievedNBooks += len(books)
		nextPage = paginationMeta.Next

		if paginationMeta.Next == 0 {
			break
		}
	}
	if retrievedNBooks != nBooks {
		t.Fatalf("Unexpected retrieved num of books. Created %v, retrieved %v", nBooks, retrievedNBooks)

	}

}
