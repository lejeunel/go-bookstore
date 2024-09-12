package controllers

import (
	"context"
	"fmt"
	m "go-bookstore/models"
	s "go-bookstore/services"

	"github.com/danielgtaylor/huma/v2"
)

type BookHTTPController struct {
	BookService s.BookService
}

func buildNestedBookRecord(b m.Book) m.BookOutputRecord {
	record := m.BookOutputRecord{Title: b.Title,
		Id: b.Id}
	for _, a := range b.Authors {
		author := m.AuthorOutputRecord{Id: a.Id, FirstName: a.FirstName, LastName: a.LastName}
		record.Authors = append(record.Authors, author)
	}

	return record

}

func (h *BookHTTPController) GetAll(ctx context.Context, pagin *m.PaginationParams) (*m.BookPaginatedOutput, error) {
	books, pagination, err := h.BookService.GetAll(ctx, *pagin)

	var out m.BookPaginatedOutput
	for _, b := range books {
		record := buildNestedBookRecord(b)
		out.Body.Data = append(out.Body.Data, record)
	}

	out.Body.Pagination = pagination

	return &out, err
}

func (h *BookHTTPController) GetOne(ctx context.Context, input *m.GetOneBookInput) (*m.BookOutput, error) {
	book, err := h.BookService.GetOne(ctx, input.Id)

	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("book with id %v not found",
			input.Id))

	}

	out := m.BookOutput{Body: buildNestedBookRecord(*book)}

	return &out, nil
}

func (h *BookHTTPController) Create(ctx context.Context, in *m.BookInput) (*m.BookOutput, error) {
	book, err := h.BookService.Create(ctx, &m.Book{Title: in.Body.Title})

	if err != nil {
		return nil, err
	}

	out := m.BookOutput{Body: buildNestedBookRecord(*book)}
	return &out, err
}

func (h *BookHTTPController) AssignAuthorToBook(ctx context.Context, in *m.AuthorBookAssign) (*m.BookOutputRecord, error) {
	book, err := h.BookService.AssignAuthorToBook(ctx, in.BookID, in.AuthorID)

	out := buildNestedBookRecord(*book)
	return &out, err
}
