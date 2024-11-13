package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	e "go-bookstore/errors"
	m "go-bookstore/models"
	"time"
)

type SQLAuthorRepo struct {
	Db *sqlx.DB
}

func NewSQLAuthorRepo(db *sqlx.DB) SQLAuthorRepo {

	return SQLAuthorRepo{Db: db}

}

func (r SQLAuthorRepo) Create(ctx context.Context, a *m.Author) (*m.Author, error) {
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

func (r SQLAuthorRepo) Delete(ctx context.Context, id string) error {

	_, err := r.Db.Exec("DELETE FROM authors WHERE id=?", id)
	return err
}

func (r SQLAuthorRepo) GetOne(ctx context.Context, id string) (*m.Author, error) {
	a := m.Author{}
	err := r.Db.Get(&a, "SELECT id,first_name,last_name,date_of_birth FROM authors WHERE id=?", id)

	if err != nil {
		return nil, err

	}

	return &a, nil
}

func (r SQLAuthorRepo) Nums() (int64, error) {
	var count int64
	err := r.Db.QueryRow("SELECT COUNT(*) FROM authors").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil

}

func (r SQLAuthorRepo) Slice(offset, length int, data interface{}) error {

	rows, err := r.Db.Queryx("SELECT id,first_name,last_name,date_of_birth FROM authors LIMIT $1 OFFSET $2", length, offset)

	if err != nil {
		return err
	}

	s := data.(*[]m.Author)

	for rows.Next() {
		var a m.Author
		err := rows.StructScan(&a)

		if err != nil {
			return err
		}

		*s = append(*s, a)
	}

	return nil
}

func (r SQLAuthorRepo) GetAuthorsOfBook(ctx context.Context, b *m.Book) ([]m.Author, error) {
	var author_ids []string
	var authors []m.Author

	err := r.Db.Select(&author_ids, "SELECT author_id FROM book_author_assoc WHERE book_id = ?", b.Id)

	if err != nil {
		return nil, e.ErrNotFound{Entity: "book", Criteria: "id", Value: b.Id.String(), Err: err}
	}

	for _, id := range author_ids {
		a, err := r.GetOne(ctx, id)
		if err != nil {
			return nil, err
		}
		authors = append(authors, *a)

	}

	return authors, nil

}
