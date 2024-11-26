package models

type PaginatedTrainsResponse struct {
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
	Total int     `json:"total"` // total number of trains
	Data  []Train `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
