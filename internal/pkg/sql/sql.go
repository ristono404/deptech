package sql

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Sanitize(s string) string {
	re, _ := regexp.Compile(`[^\w0-9A-Za-z.,\-_@_/]`)

	s = strings.TrimSpace(s)
	s = re.ReplaceAllString(s, "")

	return s
}

func Search(s string) string {
	queries := strings.Split(strings.TrimSpace(s), " ")
	newQueries := []string{}

	for _, query := range queries {
		newQueries = append(newQueries, Sanitize(query))
	}

	return strings.Join(newQueries, " ")
}

func By(s string) string {
	keys := strings.Split(strings.TrimSpace(s), ",")
	newKeys := []string{}

	for _, key := range keys {
		newKeys = append(newKeys, Sanitize(key))
	}

	return strings.Join(newKeys, ",")
}

func Params(fromDate, toDate, search, by string, isPublished bool) (string, error) {
	search = Search(search)
	by = By(by)

	result := []string{}

	if fromDate != "" {
		fromDate += "T00:00:00Z07:00"
		t, err := time.Parse(time.RFC3339, fromDate)
		if err != nil {
			return "", err
		}

		result = append(result, fmt.Sprintf("created_at >= %d", t.Unix()))
	}

	if toDate != "" {
		toDate += "T23:59:00Z07:00"
		t, err := time.Parse(time.RFC3339, toDate)
		if err != nil {
			return "", err
		}

		result = append(result, fmt.Sprintf("created_at <= %d", t.Unix()))
	}

	if search != "" && by != "" {
		for _, field := range strings.Split(by, ",") {
			result = append(result, "%s LIKE '%%s%'", field, search)
		}
	}

	if isPublished {
		result = append(result, "is_published IS TRUE")
	} else {
		result = append(result, "is_published IS FALSE OR is_published IS NULL")
	}

	return strings.Join(result, " AND "), nil
}
