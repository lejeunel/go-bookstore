package controllers

import (
	"context"
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	m "go-bookstore/models"
)

type BookHandler struct {
	Repo m.BookRepo
}

func (h *BookHandler) GetAll(ctx context.Context, input *m.PaginationParams) (*m.BookPaginatedOutput, error) {
	books, pagination, err := h.Repo.GetAll(ctx, *input)

	var out m.BookPaginatedOutput
	for _, b := range books {
		new_book := m.BookOutputRecord{Title: b.Title,
			Author: b.Author, Id: b.Id}
		out.Body.Data = append(out.Body.Data, new_book)
	}

	out.Body.Pagination = pagination

	return &out, err
}

func (h *BookHandler) GetOne(ctx context.Context, input *m.GetOneBookInput) (*m.BookOutput, error) {
	book, err := h.Repo.GetOne(ctx, input.Id)

	if err != nil {
		return nil, huma.Error404NotFound(fmt.Sprintf("book with id %v not found",
			input.Id))

	}

	out := &m.BookOutput{Body: m.BookOutputRecord{Id: book.Id, Title: book.Title,
		Author: book.Author}}

	return out, nil
}

func (h *BookHandler) Create(ctx context.Context, in *m.BookInput) (*m.BookOutput, error) {
	book, err := h.Repo.Create(ctx, &m.Book{Title: in.Body.Title,
		Author: in.Body.Author})

	out := &m.BookOutput{Body: m.BookOutputRecord{Id: book.Id, Title: book.Title,
		Author: book.Author}}
	return out, err
}
