package models

import (
	"fmt"
)

// TODO: Take this Schema from env files

var SCHEMA = "task6muxgorm"

// Book struct represents a book with its title, author, ISBN, publisher, year, genre

// TODO: DEFINE PRIMARY KEY, UNIQUE NOT NULL ETC CONSTRAINTS

// ! Update Many to many relationship between Book and User
// ? should this method be
// ** important information highlight
// * This is highlighted
// @param myParam
type Book struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	ISBN      string  `json:"isbn" gorm:"unique"`
	Publisher string  `json:"publisher"`
	Year      int     `json:"year"`
	Genre     string  `json:"genre"`
	Quantity  int     `json:"quantity" gorm:"default:0` // defaults to 0 when created
	Users     []*User `json:"users" gorm:"many2many:task6muxgorm.usersbooks;"`
}

type CreateBook struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	ISBN      string  `json:"isbn" gorm:"unique"`
	Publisher string  `json:"publisher"`
	Year      int     `json:"year"`
	Genre     string  `json:"genre"`
	Quantity  int     `json:"quantity" gorm:"default:0` // defaults to 0 when created
}

// Define Book Table Inside books Schema // TODO: How to get schema directly here from any utils package or like that.
func (b *Book) TableName() string {
	return fmt.Sprintf("%v.books", SCHEMA)
}

type User struct {
	ID    uint64  `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Books []*Book `json:"books" gorm:"many2many:task6muxgorm.usersbooks;"` // referencing to book issued // lets say one user can only issue one book
}

func (b *User) TableName() string {
	return fmt.Sprintf("%v.users", SCHEMA)
}

type BookUser struct {
	UserId uint64
	BookId uint64
}
