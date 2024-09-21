package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
	"time"
)

type SQLBookRepo struct {
	Db        *sqlx.DB
	Paginator *r.Paginator
}

func NewSQLBookRepo(db *sqlx.DB, paginator *r.Paginator) *SQLBookRepo {

	return &SQLBookRepo{Db: db, Paginator: paginator}

}

func (r *SQLBookRepo) Create(ctx context.Context, b *m.Book) (*m.Book, error) {
	b.Id = uuid.New()
	now := time.Now().String()
	query := "INSERT INTO books (id, title, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := r.Db.Exec(query, b.Id, b.Title, now,
		now)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (r *SQLBookRepo) GetOne(ctx context.Context, id string) (*m.Book, error) {
	b := m.Book{}
	err := r.Db.Get(&b, "SELECT id,title FROM books WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *SQLBookRepo) Nums() (int64, error) {
	var count int64
	err := r.Db.QueryRow("SELECT COUNT(*) FROM books ").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil

}

func (r *SQLBookRepo) Slice(offset, length int, data interface{}) error {

	rows, err := r.Db.Queryx("SELECT id, title FROM books LIMIT $1 OFFSET $2", length, offset)

	if err != nil {
		return err
	}

	var books []*m.Book

	for rows.Next() {
		var b m.Book
		err := rows.StructScan(&b)

		if err != nil {
			return err
		}

		books = append(books, &b)
	}

	data = books

	return nil
}

func (r *SQLBookRepo) AssignAuthor(ctx context.Context, b *m.Book, a *m.Author) (*m.Book, error) {
	var nAssocs int
	_ = r.Db.QueryRow("SELECT COUNT(*) FROM book_author_assoc WHERE book_id = ? AND author_id = ?",
		b.Id, a.Id).Scan(&nAssocs)

	if nAssocs > 0 {
		return b, errors.New(fmt.Sprintf("Book %v by %v is already assigned to author %v %v",
			b.Title, a.FirstName, a.LastName))
	}

	query := "INSERT INTO book_author_assoc (book_id, author_id) VALUES (?, ?)"
	_, err := r.Db.Exec(query, b.Id, a.Id)

	if err != nil {
		return nil, err
	}

	return b, nil

}
