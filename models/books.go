package models

import (
	"context"
	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title" json:"title" doc:"Title of book"`
	Author    string    `db:"author" json:"author" doc:"Name of author"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
}

type CreateBookInput struct {
	Body struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
}

type BookOutput struct {
	Body struct {
		ID     uuid.UUID `db:"id"`
		Title  string    `json:"title"`
		Author string    `json:"author"`
	}
}

type GetOneBookInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}

type BookRepo interface {
	Create(ctx context.Context, b *CreateBookInput) (*BookOutput, error)
	GetOne(ctx context.Context, input *GetOneBookInput) (*Book, error)
}
