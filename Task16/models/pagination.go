package models

type Pagination struct {
	First  int `json:"first"`
	Offset int `json:"offset"`
}
