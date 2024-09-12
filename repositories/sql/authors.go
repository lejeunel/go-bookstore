package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	m "go-bookstore/models"
	r "go-bookstore/repositories"
	"time"
)

type SQLAuthorRepo struct {
	Db        *sqlx.DB
	Paginator *r.Paginator
}

func NewSQLAuthorRepo(db *sqlx.DB, paginator *r.Paginator) *SQLAuthorRepo {

	return &SQLAuthorRepo{Db: db, Paginator: paginator}

}

func (r *SQLAuthorRepo) Create(ctx context.Context, a *m.Author) (*m.Author, error) {
	a.Id = uuid.New()
	now := time.Now().String()
	query := "INSERT INTO authors (id, first_name, last_name, date_of_birth, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.Db.Exec(query, a.Id, a.FirstName, a.LastName, a.DateOfBirth, now,
		now)

	if err != nil {
		return nil, err
	}

	return a, err
}

func (r *SQLAuthorRepo) GetOne(ctx context.Context, id string) (*m.Author, error) {
	a := m.Author{}
	err := r.Db.Get(&a, "SELECT id,first_name,last_name,date_of_birth FROM authors WHERE id=?", id)

	if err != nil {
		return nil, err

	}

	return &a, nil
}

func (r *SQLAuthorRepo) GetAll(ctx context.Context, in m.PaginationParams) ([]m.Author, *m.Pagination, error) {
	limit, offset := r.Paginator.PaginationToLimitAndOffset(in)
	authors := []m.Author{}
	err_select := r.Db.Select(&authors, "SELECT id,first_name,last_name,date_of_birth FROM authors LIMIT $1 OFFSET $2", limit, offset)
	var count int
	err_count := r.Db.Get(&count, "SELECT count(*) FROM authors")
	pagination := r.Paginator.MakePaginationMetaData(count, limit, in.Page)

	return authors, pagination, errors.Join(err_select, err_count)
}
