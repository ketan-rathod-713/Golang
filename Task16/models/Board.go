package models

type Board struct {
	ID          string `json:"Id"`
	BoardID     string `json:"boardId"`
	Visible     string `json:"visible"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}
