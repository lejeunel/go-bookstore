package controllers

import (
	"context"
	"fmt"
	g "go-bookstore/generic"
	m "go-bookstore/models"
	s "go-bookstore/services"

	"github.com/danielgtaylor/huma/v2"
)

type AuthorHTTPController struct {
	AuthorService *s.AuthorService
}

func (h *AuthorHTTPController) GetOnePage(ctx context.Context, pagin *g.PaginationParams) (*AuthorPaginatedOutput, error) {
	authors, pagination, err := h.AuthorService.GetOnePage(ctx, *pagin)

	if err != nil {
		return nil, err
	}

	var out AuthorPaginatedOutput
	for _, a := range authors {
		out.Body.Data = append(out.Body.Data, AuthorOutputRecord{Id: a.Id, FirstName: a.FirstName,
			LastName: a.LastName, DateOfBirth: a.DateOfBirth})
	}

	out.Body.Pagination = pagination

	return &out, nil
}

func (h *AuthorHTTPController) GetOne(ctx context.Context, input *GetOneAuthorInput) (*AuthorOutput, error) {
	a, err := h.AuthorService.GetOne(ctx, input.Id)

	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("author with id %v not found",
			input.Id))

	}

	out := AuthorOutput{Body: AuthorOutputRecord{Id: a.Id, FirstName: a.FirstName,
		LastName: a.LastName, DateOfBirth: a.DateOfBirth}}

	return &out, nil
}

func (h *AuthorHTTPController) Create(ctx context.Context, in *AuthorMessage) (*AuthorOutput, error) {
	a, err := h.AuthorService.Create(ctx, &m.Author{FirstName: in.Body.FirstName,
		LastName: in.Body.LastName, DateOfBirth: in.Body.DateOfBirth})

	if err != nil {
		return nil, err
	}
	out := AuthorOutput{Body: AuthorOutputRecord{Id: a.Id, FirstName: a.FirstName,
		LastName: a.LastName, DateOfBirth: a.DateOfBirth}}

	return &out, err
}
