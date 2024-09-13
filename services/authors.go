package services

import (
	"context"
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

func (s *AuthorService) GetAll(ctx context.Context, in m.PaginationParams) ([]m.Author, *m.Pagination, error) {
	authors, pagination, err := s.AuthorRepo.GetAll(ctx, in)

	return authors, pagination, err
}
