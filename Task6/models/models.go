package models

const SCHEMA = "task6MuxGorm"

// Book struct represents a book with its title, author, ISBN, publisher, year, genre
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Year      int
	Genre     string
}

// Define Book Table Inside books Schema
func (b Book) TableName() string {
	return SCHEMA + ".books"
}
