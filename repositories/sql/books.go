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
	err := r.Db.Get(&b, "SELECT * FROM books WHERE id=$1", id)

	return &b, err
}

func (r *SQLBookRepo) GetAll(ctx context.Context, in m.PaginationParams) ([]m.Book, *m.Pagination, error) {
	limit, offset := r.Paginator.PaginationToLimitAndOffset(in)
	books := []m.Book{}
	rows, err := r.Db.Queryx("SELECT id, title FROM books LIMIT $1 OFFSET $2", limit, offset)
	for rows.Next() {
		var b m.Book
		err := rows.StructScan(&b)

		if err != nil {
			return nil, nil, err
		}

		books = append(books, b)
	}
	pagination := r.Paginator.MakePaginationMetaData(len(books), limit, in.Page)

	return books, pagination, err
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
