package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	m "go-bookstore/models"
	"time"
)

var book_schema = `
CREATE TABLE IF NOT EXISTS books (
    id varchar(16),
    title text,
    author text,
	created_at text,
	updated_at text
);`

type BookRepo struct {
	Db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) BookRepo {

	db.MustExec(book_schema)

	return BookRepo{Db: db}

}

func (r *BookRepo) Create(ctx context.Context, in *m.CreateBookInput) (*m.BookOutput, error) {
	id := uuid.New()
	now := time.Now().String()
	query := "INSERT INTO books (id, title, author, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.Db.Exec(query, id, in.Body.Title, in.Body.Author, now,
		now)

	out := &m.BookOutput{}
	out.Body.ID = id
	out.Body.Title = in.Body.Title
	out.Body.Author = in.Body.Author

	return out, err
}

func (r *BookRepo) GetOne(ctx context.Context, id string) (*m.Book, error) {
	b := m.Book{}
	err := r.Db.Get(&b, "SELECT * FROM books WHERE id=$1", id)

	return &b, err
}
