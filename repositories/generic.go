package repositories

import (
	"context"
	m "go-bookstore/models"
)

type BookRepo interface {
	Create(context.Context, *m.Book) (*m.Book, error)
	GetOne(context.Context, string) (*m.Book, error)
	AssignAuthor(context.Context, *m.Book, *m.Author) (*m.Book, error)
	Nums() (int64, error)
	Slice(offset, length int, data interface{}) error
}

type AuthorRepo interface {
	Create(context.Context, *m.Author) (*m.Author, error)
	GetOne(context.Context, string) (*m.Author, error)
	GetAuthorsOfBook(context.Context, *m.Book) ([]m.Author, error)
	Nums() (int64, error)
	Slice(offset, length int, data interface{}) error
}
