package repositories

import (
	m "go-bookstore/models"
)

type Paginator struct {
	MaxPageSize int
}

func (p *Paginator) PaginationToLimitAndOffset(params m.PaginationParams) (int, int) {
	limit := max(p.MaxPageSize, params.PageSize)
	offset := (params.Page - 1) * limit

	return limit, offset
}

func (p *Paginator) MakePaginationMetaData(recordCount int, limit, page int) *m.Pagination {
	res := m.Pagination{}

	total := (recordCount / limit)

	// Calculator Total Page
	remainder := (recordCount % limit)
	if remainder == 0 {
		res.TotalPage = total
	} else {
		res.TotalPage = total + 1
	}

	// Set current/record per page meta data
	res.CurrentPage = page
	res.RecordPerPage = limit

	// Calculator the Next/Previous Page
	if page <= 0 {
		res.Next = page + 1
	} else if page < res.TotalPage {
		res.Previous = page - 1
		res.Next = page + 1
	} else if page == res.TotalPage {
		res.Previous = page - 1
		res.Next = 0
	}

	return &res
}
