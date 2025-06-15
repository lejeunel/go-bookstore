package services

import (
	"context"
	e "go-bookstore/errors"
	g "go-bookstore/generic"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type AuthorService struct {
	AuthorRepo      *r.AuthorRepo
	BookRepo        *r.BookRepo
	MaxPageSize     int
	DefaultPageSize int
}

func NewAuthorService(authorRepo r.AuthorRepo, bookRepo r.BookRepo, maxPageSize, defaultPageSize int) *AuthorService {
	return &AuthorService{AuthorRepo: &authorRepo, BookRepo: &bookRepo, MaxPageSize: maxPageSize, DefaultPageSize: defaultPageSize}
}

func (s *AuthorService) Create(ctx context.Context, a *m.Author) (*m.Author, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	return (*s.AuthorRepo).Create(ctx, a)
}

func (s *AuthorService) GetBooksOfAuthor(ctx context.Context, author *m.Author) ([]m.Book, error) {
	books, err := (*s.BookRepo).GetBooksOfAuthor(ctx, author)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *AuthorService) Delete(ctx context.Context, id string) error {
	author, err := (*s.AuthorRepo).GetOne(ctx, id)
	if err != nil {
		return err
	}

	books, err := (*s.BookRepo).GetBooksOfAuthor(ctx, author)
	if len(books) > 0 {
		return e.ErrForbiddenDeletingDependency{ParentEntity: "author", ParentId: author.Id.String(), ChildEntity: "book"}
	}

	return (*s.AuthorRepo).Delete(ctx, id)
}

func (s *AuthorService) GetOne(ctx context.Context, id string) (*m.Author, error) {
	return (*s.AuthorRepo).GetOne(ctx, id)
}

func (s *AuthorService) GetOnePage(ctx context.Context, in g.PaginationParams) ([]m.Author, *g.PaginationMeta, error) {
	p, err := g.NewPaginator(*s.AuthorRepo, in.PageSize, s.MaxPageSize, in.Page)
	if err != nil {
		return nil, nil, err
	}
	var authors []m.Author

	if err := p.Results(&authors); err != nil {
		return nil, nil, err
	}

	paginationMeta := g.NewPaginationMeta(p)

	return authors, &paginationMeta, nil

}
