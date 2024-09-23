package services

import (
	"errors"
	"fmt"
	pag "github.com/vcraescu/go-paginator/v2"
)

func NewPaginator(adapter pag.Adapter, pageSize, maxPerPage, page int) (pag.Paginator, error) {

	if pageSize > maxPerPage {
		return nil, errors.New(fmt.Sprintf("Provided page size %d must be <= to %d", pageSize, maxPerPage))
	}

	if page < 1 {
		return nil, errors.New(fmt.Sprintf("Provided page %d must be >= to 1", page))
	}

	p := pag.New(adapter, maxPerPage)
	p.SetPage(page)
	return p, nil
}
