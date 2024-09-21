package services

import (
	"context"
	"github.com/vcraescu/go-paginator/v2"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type AuthorService struct {
	AuthorRepo r.AuthorRepo
}

func (s *AuthorService) Create(ctx context.Context, a *m.Author) (*m.Author, error) {
	return s.AuthorRepo.Create(ctx, a)
}

func (s *AuthorService) GetOne(ctx context.Context, id string) (*m.Author, error) {
	return s.AuthorRepo.GetOne(ctx, id)
}

func (s *AuthorService) GetMany(ctx context.Context, in m.PaginationParams) ([]m.Author, *m.Pagination, error) {
	p := paginator.New(s.AuthorRepo, in.PageSize)
	p.SetPage(in.Page)
	var authors []m.Author

	if err := p.Results(&authors); err != nil {
		return nil, nil, err
	}

	return authors, nil, nil

}
