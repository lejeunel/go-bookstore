package config

import (
	"github.com/jmoiron/sqlx"
	c "go-bookstore/controllers"
	r "go-bookstore/repositories"
	sql "go-bookstore/repositories/sql"
	s "go-bookstore/services"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddRoutes(api huma.API, db *sqlx.DB, p *r.Paginator, prefix string) {
	bookService := s.BookService{BookRepo: sql.NewSQLBookRepo(db, p)}
	bookController := &c.BookHTTPController{BookService: bookService}

	authorService := s.AuthorService{AuthorRepo: sql.NewSQLAuthorRepo(db, p)}
	authorController := &c.AuthorHTTPController{AuthorService: authorService}

	huma.Register(api, huma.Operation{
		OperationID: "get-book",
		Method:      http.MethodGet,
		Path:        prefix + "/books/{id}",
		Summary:     "Get a book by ID",
	}, bookController.GetOne)
	huma.Register(api, huma.Operation{
		OperationID: "get-all-books",
		Method:      http.MethodGet,
		Path:        prefix + "/books",
		Summary:     "Get all books",
	}, bookController.GetAll)

	huma.Register(api, huma.Operation{
		OperationID:   "post-book",
		Method:        http.MethodPost,
		Path:          prefix + "/books",
		Summary:       "Add a new book",
		DefaultStatus: http.StatusCreated,
	}, bookController.Create)

	huma.Register(api, huma.Operation{
		OperationID:   "assign-author",
		Method:        http.MethodPost,
		Path:          prefix + "/books/{book_id}/authors/{author_id}",
		Summary:       "Assign an author to book",
		DefaultStatus: http.StatusCreated,
	}, bookController.AssignAuthorToBook)

	huma.Register(api, huma.Operation{
		OperationID:   "post-author",
		Method:        http.MethodPost,
		Path:          prefix + "/authors",
		Summary:       "Add a new author",
		DefaultStatus: http.StatusCreated,
	}, authorController.Create)
	huma.Register(api, huma.Operation{
		OperationID: "get-all-authors",
		Method:      http.MethodGet,
		Path:        prefix + "/authors",
		Summary:     "Get all authors",
	}, authorController.GetAll)
}
