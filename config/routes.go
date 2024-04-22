package config

import (
	"github.com/jmoiron/sqlx"
	"go-bookstore/controllers"
	r "go-bookstore/repositories"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddRoutes(api huma.API, db *sqlx.DB, p *r.Paginator, prefix string) {
	bookHandler := controllers.NewSQLBookHandler(db, p)

	// TODO add prefix to all paths here...

	// v1.Handle("/book/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods(http.MethodDelete)
	huma.Register(api, huma.Operation{
		OperationID: "get-book",
		Method:      http.MethodGet,
		Path:        prefix + "/books/{id}",
		Summary:     "Get a book by ID",
	}, bookHandler.GetOne)
	huma.Register(api, huma.Operation{
		OperationID: "get-all-book",
		Method:      http.MethodGet,
		Path:        prefix + "/books",
		Summary:     "Get all books",
	}, bookHandler.GetAll)

	huma.Register(api, huma.Operation{
		OperationID:   "post-book",
		Method:        http.MethodPost,
		Path:          prefix + "/books",
		Summary:       "Add a new book",
		DefaultStatus: http.StatusCreated,
	}, bookHandler.Create)
}
