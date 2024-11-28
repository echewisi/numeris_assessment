package utils

import (
	"net/http"
	"strconv"
)

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Limit  int
	Offset int
}

// ParsePagination extracts pagination parameters from the query string
func ParsePagination(r *http.Request) PaginationParams {
	query := r.URL.Query()

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil || offset < 0 {
		offset = 0 // Default offset
	}

	return PaginationParams{
		Limit:  limit,
		Offset: offset,
	}
}

// PaginateSlice paginates a slice based on the pagination parameters
func PaginateSlice[T any](data []T, params PaginationParams) []T {
	start := params.Offset
	end := start + params.Limit

	if start > len(data) {
		return []T{}
	}

	if end > len(data) {
		end = len(data)
	}

	return data[start:end]
}
