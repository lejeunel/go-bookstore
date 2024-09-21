package models

import paginator "github.com/vcraescu/go-paginator/v2"

type PaginationParams struct {
	Page     int `query:"page" minimum:"1" default:"1"`
	PageSize int `query:"pagesize"`
}

type Pagination struct {
	Next         int   `json:"next,omitempty"`
	Previous     int   `json:"prev,omitempty"`
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

func NewPagination(p paginator.Paginator) Pagination {

	pagination := Pagination{}

	hasNext, _ := p.HasNext()

	if hasNext {
		pagination.Next, _ = p.NextPage()
	}

	hasPrev, _ := p.HasPrev()

	if hasPrev {
		pagination.Previous, _ = p.PrevPage()
	}

	pagination.CurrentPage, _ = p.Page()
	pagination.TotalPages, _ = p.PageNums()
	pagination.TotalRecords, _ = p.Nums()

	return pagination

}
