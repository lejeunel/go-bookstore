package models

type PaginationParams struct {
	Page     int `query:"page" minimum:"1" default:"1"`
	PageSize int `query:"pagesize"`
}

type Pagination struct {
	Next          int `json:"next,omitempty"`
	Previous      int `json:"prev,omitempty" minimum:"1"`
	RecordPerPage int `json:"records_per_page"`
	CurrentPage   int `json:"current_page"`
	TotalPage     int `json:"total_pages"`
}
