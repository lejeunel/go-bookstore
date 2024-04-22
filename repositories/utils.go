package repositories

import (
	m "go-bookstore/models"
)

type Paginator struct {
	MaxPageSize int
}

func (p *Paginator) pagination_to_limit_offset(params m.PaginationParams) (int, int) {
	limit := max(p.MaxPageSize, params.PageSize)
	offset := (params.Page - 1) * limit

	return limit, offset
}
