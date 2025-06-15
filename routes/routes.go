package config

import (
	c "go-bookstore/controllers"
	s "go-bookstore/services"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddAPIRoutes(api huma.API, prefix string, bookService s.BookService,
	authorService s.AuthorService) {

	bookController := &c.BookHTTPController{BookService: &bookService}
	authorController := &c.AuthorHTTPController{AuthorService: &authorService}

	huma.Register(api, huma.Operation{
		OperationID: "get-book",
		Method:      http.MethodGet,
		Path:        prefix + "/books/{id}",
		Tags:        []string{"Books"},
		Summary:     "Get a book by ID",
	}, bookController.GetOne)
	huma.Register(api, huma.Operation{
		OperationID: "get-many-books",
		Method:      http.MethodGet,
		Path:        prefix + "/books",
		Tags:        []string{"Books"},
		Summary:     "Get several books",
	}, bookController.GetOnePage)

	huma.Register(api, huma.Operation{
		OperationID:   "post-book",
		Method:        http.MethodPost,
		Path:          prefix + "/books",
		Tags:          []string{"Books"},
		Summary:       "Add a new book",
		DefaultStatus: http.StatusCreated,
	}, bookController.Create)

	huma.Register(api, huma.Operation{
		OperationID:   "assign-author",
		Method:        http.MethodPost,
		Path:          prefix + "/books/{book_id}/authors/{author_id}",
		Summary:       "Assign an author to book",
		Tags:          []string{"Books"},
		DefaultStatus: http.StatusCreated,
	}, bookController.AssignAuthorToBook)

	huma.Register(api, huma.Operation{
		OperationID:   "post-author",
		Method:        http.MethodPost,
		Path:          prefix + "/authors",
		Summary:       "Add a new author",
		Tags:          []string{"Authors"},
		DefaultStatus: http.StatusCreated,
	}, authorController.Create)
	huma.Register(api, huma.Operation{
		OperationID: "get-many-authors",
		Method:      http.MethodGet,
		Path:        prefix + "/authors",
		Tags:        []string{"Authors"},
		Summary:     "Get many authors",
	}, authorController.GetOnePage)

	huma.Register(api, huma.Operation{
		OperationID: "get-author",
		Method:      http.MethodGet,
		Path:        prefix + "/authors/{id}",
		Tags:        []string{"Authors"},
		Summary:     "Get an author by id",
	}, authorController.GetOne)
}
