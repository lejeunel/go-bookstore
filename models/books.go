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

type BookOutputBody struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type BookOutput struct {
	Body BookOutputBody
}

type BookOutputList struct {
	Body []BookOutputBody
}

type GetOneBookInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}

type BookRepo interface {
	Create(ctx context.Context, b *Book) (*Book, error)
	GetOne(ctx context.Context, id string) (*Book, error)
	GetAll(ctx context.Context, id PaginationParams) ([]Book, error)
}
