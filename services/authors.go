package services

import (
	"context"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type AuthorService struct {
	AuthorRepo      r.AuthorRepo
	MaxPageSize     int
	DefaultPageSize int
}

func (s *AuthorService) Create(ctx context.Context, a *m.Author) (*m.Author, error) {
	return s.AuthorRepo.Create(ctx, a)
}

func (s *AuthorService) GetOne(ctx context.Context, id string) (*m.Author, error) {
	return s.AuthorRepo.GetOne(ctx, id)
}

func (s *AuthorService) GetMany(ctx context.Context, in m.PaginationParams) ([]m.Author, *m.PaginationMeta, error) {
	p, err := NewPaginator(s.AuthorRepo, in.PageSize, s.MaxPageSize, in.Page)
	if err != nil {
		return nil, nil, err
	}
	var authors []m.Author

	if err := p.Results(&authors); err != nil {
		return nil, nil, err
	}

	paginationMeta := m.NewPaginationMeta(p)

	return authors, &paginationMeta, nil

}
