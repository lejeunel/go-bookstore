package models

import (
	"github.com/google/uuid"
)

type Author struct {
	Id          uuid.UUID
	FirstName   string
	LastName    string
	DateOfBirth string
}

type AuthorInputBody struct {
	FirstName   string `json:"first_name" doc:"First name"`
	LastName    string `json:"last_name" doc:"Last name"`
	DateOfBirth string `json:"date_of_birth" doc:"Date of birth"`
}

type AuthorInput struct {
	Body AuthorInputBody
}

type AuthorOutputRecord struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth" doc:"Date of birth"`
}

type AuthorOutput struct {
	Body AuthorOutputRecord
}

type AuthorPaginatedOutput struct {
	Body struct {
		Pagination *Pagination          `json:"pagination"`
		Data       []AuthorOutputRecord `json:"data"`
	}
}

type GetOneAuthorInput struct {
	Id string `path:"id" doc:"ID to retrieve"`
}
