package models

import (
	"github.com/google/uuid"
)

type Author struct {
	Id          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	DateOfBirth string    `db:"date_of_birth"`
}

type AuthorInputBody struct {
	FirstName   string `json:"first_name" doc:"First name"`
	LastName    string `json:"last_name" doc:"Last name"`
	DateOfBirth string `json:"date_of_birth,omitempty" doc:"Date of birth"`
}

type AuthorInput struct {
	Body AuthorInputBody
}

type AuthorOutputRecord struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth,omitempty"`
}

type AuthorOutput struct {
	Body AuthorOutputRecord
}

type AuthorPaginatedOutputBody struct {
	Pagination *PaginationMeta      `json:"pagination"`
	Data       []AuthorOutputRecord `json:"data"`
}

type AuthorPaginatedOutput struct {
	Body AuthorPaginatedOutputBody
}

type GetOneAuthorInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}
