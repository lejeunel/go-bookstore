package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/vcraescu/go-paginator/v2"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type BookService struct {
	BookRepo        r.BookRepo
	AuthorRepo      r.AuthorRepo
	MaxPageSize     int
	DefaultPageSize int
}

func (s *BookService) Create(ctx context.Context, b *m.Book) (*m.Book, error) {

	return s.BookRepo.Create(ctx, b)
}

func (s *BookService) GetOne(ctx context.Context, id string) (*m.Book, error) {
	book, err := s.BookRepo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	authors, err := s.AuthorRepo.GetAuthorsOfBook(ctx, book)

	if err != nil {
		return nil, err
	}

	book.Authors = authors

	return book, nil

}

func (s *BookService) GetMany(ctx context.Context, in m.PaginationParams) ([]m.Book, *m.Pagination, error) {
	p := paginator.New(s.BookRepo, s.MaxPageSize)
	if in.Page == 0 {
		p.SetPage(1)
	} else {
		p.SetPage(in.Page)
	}

	if in.PageSize > s.MaxPageSize {
		return nil, nil, errors.New(fmt.Sprintf("Provided page size %d must be <= to %d", in.PageSize, s.MaxPageSize))
	}

	var books []m.Book
	err := p.Results(&books)

	if err != nil {
		return nil, nil, err
	}

	pagination := m.NewPagination(p)
	return books, &pagination, nil

}

func (s *BookService) AssignAuthorToBook(ctx context.Context, book_id string, author_id string) (*m.Book, error) {
	book, err_book := s.GetOne(ctx, book_id)
	author, err_author := s.AuthorRepo.GetOne(ctx, author_id)

	if (err_book != nil) || (err_author != nil) {
		return nil, errors.Join(err_book, err_author)
	}

	book, err := s.BookRepo.AssignAuthor(ctx, book, author)
	if err != nil {
		return nil, err
	}

	book.Authors = append(book.Authors, *author)

	return book, nil
}
