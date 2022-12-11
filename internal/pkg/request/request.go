package request

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/deptech/internal/pkg/sql"
)

func List(r url.Values, defaultPerPage uint64) (page, limit, offset uint64, from, to *time.Time, err error) {
	page = 1
	if r.Get("page") != "" {
		if page, err = strconv.ParseUint(r.Get("page"), 10, 64); err != nil {
			return
		}
	}

	limit = defaultPerPage
	if r.Get("per_page") != "" {
		if limit, err = strconv.ParseUint(r.Get("per_page"), 10, 64); err != nil {
			return
		}
	}

	if page > 0 {
		offset = (page - 1) * limit
	}

	fromParam := r.Get("from")
	if fromParam != "" && len(fromParam) == 10 {
		var dateParam time.Time
		format := fmt.Sprintf("%s 00:00:00", fromParam)
		dateParam, err = time.Parse("2006-01-02 15:04:05", format)
		if err != nil {
			return
		}
		from = &dateParam
	}

	toParam := r.Get("to")
	if toParam != "" && len(toParam) == 10 {
		var dateParam time.Time
		format := fmt.Sprintf("%s 23:59:59", toParam)
		dateParam, err = time.Parse("2006-01-02 15:04:05", format)
		if err != nil {
			return
		}
		to = &dateParam
	}

	return
}

func Params(r *http.Request, defaultPerPage int) (page, start, perpage int, is_published *int, searchReq Search) {
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	start, _ = strconv.Atoi(r.URL.Query().Get("start"))
	perpage, _ = strconv.Atoi(r.URL.Query().Get("per_page"))
	if perpage <= 0 {
		perpage = defaultPerPage
	}
	if page <= 0 {
		page = 1
	}
	searchReq = Search{
		By:       sql.By(r.URL.Query().Get("by")),
		Search:   sql.Search(r.URL.Query().Get("search")),
		FromDate: sql.Sanitize(r.URL.Query().Get("from")),
		ToDate:   sql.Sanitize(r.URL.Query().Get("to")),
	}
	publish := r.URL.Query().Get("is_published")
	if publish != "" {
		val, _ := strconv.Atoi(publish)
		is_published = &val
	}

	return
}
