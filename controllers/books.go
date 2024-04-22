package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/danielgtaylor/huma/v2"
	"github.com/jmoiron/sqlx"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
)

type BookHandler struct {
	repo m.BookRepo
}

func NewSQLBookHandler(db *sqlx.DB, p *r.Paginator) *BookHandler {
	return &BookHandler{repo: r.NewSQLBookRepo(db, p)}
}

func (h *BookHandler) GetAll(ctx context.Context, input *m.PaginationParams) (*m.BookOutputList, error) {
	books, _ := h.repo.GetAll(ctx, *input)

	var out m.BookOutputList
	for _, b := range books {
		new_book := m.BookOutputBody{Title: b.Title,
			Author: b.Author, Id: b.Id}
		out.Body = append(out.Body, new_book)
	}

	return &out, nil
}

func (h *BookHandler) GetOne(ctx context.Context, input *m.GetOneBookInput) (*m.BookOutput, error) {
	book, err := h.repo.GetOne(ctx, input.Id)

	if err == sql.ErrNoRows {
		return nil, huma.Error404NotFound(fmt.Sprintf("book with id %v not found",
			input.Id))

	}

	out := &m.BookOutput{Body: m.BookOutputBody{Id: book.Id, Title: book.Title,
		Author: book.Author}}

	return out, nil
}

func (h *BookHandler) Create(ctx context.Context, in *m.BookInput) (*m.BookOutput, error) {
	book, err := h.repo.Create(ctx, &m.Book{Title: in.Body.Title,
		Author: in.Body.Author})

	out := &m.BookOutput{Body: m.BookOutputBody{Id: book.Id, Title: book.Title,
		Author: book.Author}}
	return out, err
}
