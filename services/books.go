package services

import (
	"context"
	"errors"
	"github.com/vcraescu/go-paginator/v2"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type BookService struct {
	BookRepo   r.BookRepo
	AuthorRepo r.AuthorRepo
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
	p := paginator.New(s.BookRepo, in.PageSize)
	p.SetPage(in.Page)
	var books []m.Book

	if err := p.Results(&books); err != nil {
		return nil, nil, err
	}

	return books, nil, nil

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
