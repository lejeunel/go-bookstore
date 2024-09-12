package repositories

import (
	"context"
	m "go-bookstore/models"
)

type BookRepo interface {
	Create(ctx context.Context, b *m.Book) (*m.Book, error)
	GetOne(ctx context.Context, id string) (*m.Book, error)
	GetAll(ctx context.Context, id m.PaginationParams) ([]m.Book, *m.Pagination, error)
	AssignAuthor(ctx context.Context, b *m.Book, a *m.Author) (*m.Book, error)
}

type AuthorRepo interface {
	Create(ctx context.Context, b *m.Author) (*m.Author, error)
	GetOne(ctx context.Context, id string) (*m.Author, error)
	GetAll(ctx context.Context, id m.PaginationParams) ([]m.Author, *m.Pagination, error)
}
