package models

import "gorm.io/gorm"

const SCHEMA = "task6muxgorm"

// Book struct represents a book with its title, author, ISBN, publisher, year, genre

// TODO: DEFINE PRIMARY KEY, UNIQUE NOT NULL ETC CONSTRAINTS
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
}

// Define Book Table Inside books Schema // TODO: How to get schema directly here from any utils package or like that.
func (b Book) TableName() string {
	return SCHEMA + ".books"
}
