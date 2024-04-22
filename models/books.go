package models

import (
	"context"
	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
}

type BookInputBody struct {
	Title  string `json:"title" doc:"Title of book"`
	Author string `json:"author" doc:"Name of author"`
}

type BookInput struct {
	Body BookInputBody
}

type BookOutputRecord struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type BookOutput struct {
	Body BookOutputRecord
}

type BookPaginatedOutput struct {
	Body struct {
		Pagination *Pagination        `json:"pagination"`
		Data       []BookOutputRecord `json:"data"`
	}
}

type GetOneBookInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}

type BookRepo interface {
	Create(ctx context.Context, b *Book) (*Book, error)
	GetOne(ctx context.Context, id string) (*Book, error)
	GetAll(ctx context.Context, id PaginationParams) ([]Book, *Pagination, error)
}
