package models

import paginator "github.com/vcraescu/go-paginator/v2"

type PaginationParams struct {
	Page     int `query:"page" minimum:"1" default:"1"`
	PageSize int `query:"pagesize"`
}

type Pagination struct {
	Next        int `json:"next"`
	Previous    int `json:"prev"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_pages"`
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
	pagination.TotalPage, _ = p.PageNums()

	return pagination

}
