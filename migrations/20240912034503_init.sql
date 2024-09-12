-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS authors (
    id varchar(16),
    first_name text,
    last_name text,
    date_of_birth text,
	created_at text,
	updated_at text
);


CREATE TABLE IF NOT EXISTS books (
    id varchar(16),
    title text,
    created_at text,
    updated_at text
);
CREATE TABLE IF NOT EXISTS book_author_assoc (
    book_id varchar(16),
    author_id varchar(16)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE authors;
DROP TABLE books;

-- +goose StatementEnd
