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
	repo r.BookRepo
}

func NewBookHandler(db *sqlx.DB) *BookHandler {
	return &BookHandler{
		repo: r.NewBookRepo(db),
	}
}

func (h *BookHandler) GetOne(ctx context.Context, input *m.GetOneBookInput) (*m.BookOutput, error) {
	book, err := h.repo.GetOne(ctx, input.Id)

	if err == sql.ErrNoRows {
		return nil, huma.Error404NotFound(fmt.Sprintf("book with id %v not found",
			input.Id))

	}

	out := &m.BookOutput{}
	out.Body.ID = book.ID
	out.Body.Title = book.Title
	out.Body.Author = book.Author

	return out, nil
}

func (h *BookHandler) Create(ctx context.Context, in *m.CreateBookInput) (*m.BookOutput, error) {
	book, err := h.repo.Create(ctx, &m.Book{Title: in.Body.Title,
		Author: in.Body.Author})

	out := &m.BookOutput{}
	out.Body.ID = book.ID
	out.Body.Title = book.Title
	out.Body.Author = book.Author
	return out, err
}
