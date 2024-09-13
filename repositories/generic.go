package repositories

import (
	"context"
	m "go-bookstore/models"
)

type BookRepo interface {
	Create(context.Context, *m.Book) (*m.Book, error)
	GetOne(context.Context, string) (*m.Book, error)
	GetAll(context.Context, m.PaginationParams) ([]m.Book, *m.Pagination, error)
	AssignAuthor(context.Context, *m.Book, *m.Author) (*m.Book, error)
}

type AuthorRepo interface {
	Create(context.Context, *m.Author) (*m.Author, error)
	GetOne(context.Context, string) (*m.Author, error)
	GetAll(context.Context, m.PaginationParams) ([]m.Author, *m.Pagination, error)
	GetAuthorsOfBook(context.Context, *m.Book) ([]m.Author, error)
}
