package services

import (
	"context"
	"errors"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type BookService struct {
	BookRepo   r.BookRepo
	AuthorRepo r.AuthorRepo
}

func (s *BookService) Create(ctx context.Context, b *m.Book) (*m.Book, error) {

	b, err := s.BookRepo.Create(ctx, b)

	return b, err
}

func (s *BookService) GetOne(ctx context.Context, id string) (*m.Book, error) {
	book, err := s.BookRepo.GetOne(ctx, id)

	return book, err
}

func (s *BookService) GetAll(ctx context.Context, in m.PaginationParams) ([]m.Book, *m.Pagination, error) {
	books, pagination, error := s.BookRepo.GetAll(ctx, in)

	return books, pagination, error
}

func (s *BookService) AssignAuthorToBook(ctx context.Context, book_id string, author_id string) (*m.Book, error) {
	book, err_book := s.BookRepo.GetOne(ctx, book_id)
	author, err_author := s.AuthorRepo.GetOne(ctx, author_id)

	if (err_book != nil) || (err_author != nil) {
		return nil, errors.Join(err_book, err_author)
	}

	book, err := s.BookRepo.AssignAuthor(ctx, book, author)
	if err != nil {
		return nil, err
	}

	return book, nil
}
