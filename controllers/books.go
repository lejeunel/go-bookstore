package controllers

import (
	"context"
	g "go-bookstore/generic"
	m "go-bookstore/models"
	s "go-bookstore/services"

	"github.com/danielgtaylor/huma/v2"
)

type BookHTTPController struct {
	BookService *s.BookService
}

func buildNestedBookRecord(b m.Book) BookOutputRecord {
	record := BookOutputRecord{Title: b.Title,
		Id: b.Id}
	for _, a := range b.Authors {
		author := AuthorOutputRecord{Id: a.Id, FirstName: a.FirstName, LastName: a.LastName}
		record.Authors = append(record.Authors, author)
	}

	return record

}

func (h *BookHTTPController) GetMany(ctx context.Context, pagin *g.PaginationParams) (*BookPaginatedOutput, error) {
	books, pagination, err := h.BookService.GetMany(ctx, *pagin)

	var out BookPaginatedOutput
	for _, b := range books {
		record := buildNestedBookRecord(b)
		out.Body.Data = append(out.Body.Data, record)
	}

	out.Body.Pagination = pagination

	return &out, err
}

func (h *BookHTTPController) GetOne(ctx context.Context, input *GetOneBookInput) (*BookOutput, error) {
	book, err := h.BookService.GetOne(ctx, input.Id)

	if err != nil {
		return nil, huma.Error404NotFound(err.Error(), err)

	}

	out := BookOutput{Body: buildNestedBookRecord(*book)}

	return &out, nil
}

func (h *BookHTTPController) Create(ctx context.Context, in *BookMessage) (*BookOutput, error) {
	book, err := h.BookService.Create(ctx, &m.Book{Title: in.Body.Title})

	if err != nil {
		return nil, err
	}

	out := BookOutput{Body: buildNestedBookRecord(*book)}
	return &out, err
}

func (h *BookHTTPController) AssignAuthorToBook(ctx context.Context, in *AuthorBookAssignInput) (*BookOutput, error) {
	book, err := h.BookService.AssignAuthorToBook(ctx, in.BookID, in.AuthorID)

	if err != nil {
		return nil, err
	}

	out := BookOutput{Body: buildNestedBookRecord(*book)}
	return &out, err
}
