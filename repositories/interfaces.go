package repositories

import (
	c "context"
	m "go-bookstore/models"
)

type BookRepo interface {
	Create(c.Context, *m.Book) (*m.Book, error)
	Delete(c.Context, string) error
	Find(c.Context, string) (*m.Book, error)
	AssignAuthor(c.Context, *m.Book, *m.Author) (*m.Book, error)
	GetBooksOfAuthor(c.Context, *m.Author) ([]m.Book, error)
	Nums() (int64, error)
	Slice(offset, length int, data interface{}) error
}

type AuthorRepo interface {
	Create(c.Context, *m.Author) (*m.Author, error)
	Delete(c.Context, string) error
	GetOne(c.Context, string) (*m.Author, error)
	GetAuthorsOfBook(c.Context, *m.Book) ([]m.Author, error)
	Nums() (int64, error)
	Slice(offset, length int, data interface{}) error
}
