package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Author struct {
	Id          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	DateOfBirth string    `db:"date_of_birth"`
}

func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.FirstName, validation.Required),
		validation.Field(&a.LastName, validation.Required),
		validation.Field(&a.DateOfBirth, validation.Date("2006-01-02")),
	)
}
