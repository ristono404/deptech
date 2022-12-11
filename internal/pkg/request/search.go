package request

import (
	"github.com/ristono404/deptech/internal/pkg/sql"
)

type Search struct {
	Search   string `json:"search"`
	By       string `json:"by"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

func (s *Search) Query(isPublished bool) (string, error) {
	return sql.Params(s.FromDate, s.ToDate, s.FromDate, s.ToDate, isPublished)
}

func New(search, by, fromDate, toDate string) *Search {
	return &Search{
		Search:   sql.Search(search),
		By:       sql.By(by),
		FromDate: sql.Sanitize(fromDate),
		ToDate:   sql.Sanitize(toDate),
	}
}
