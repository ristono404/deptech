package response

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Pages struct {
	Total      uint64 `json:"total"`
	PerPage    uint64 `json:"per_page"`
	Current    uint64 `json:"current"`
	TotalPages uint64 `json:"total_pages"`
}

type View struct {
	Status       bool        `json:"status"`
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
	Pagination   Pages       `json:"pagination"`
}

func Response(w http.ResponseWriter, view interface{}, httpStatus int) error {
	res, err := json.Marshal(view)
	if err != nil {
		return err
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, x-api-Key, X-localization, channel, Channel, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}
