package bookservice

/*
	Book Services
	- create a new book
	- get all books
	- get one book by id
	- update one book
	- delete one book
*/

import (
	"errors"
	"log"
	"task6MuxGorm/app"
	"task6MuxGorm/models"

	"gorm.io/gorm"
)

/*
define all services a bookservice needs.
*/
type Service interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBooks() ([]*models.Book, error)
	GetOneBookById(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id uint) (*models.Book, error)
}

/*
instance of db with required data which needed for performing its services.
here only database connection required, as it will be doing only raw operations.
*/
type service struct { // Implement Service Interface
	DB *gorm.DB
}

/*
create new book service
get DB from app and initialise it.
*/
func New(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// TODO: Implement Service Interface

/*
for given data, insert raw in book table.
*/
func (s *service) CreateBook(book *models.Book) (*models.Book, error) {
	result := s.DB.Create(book)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	log.Println("Book inserted", *book)
	return book, nil
}

/*
Get all the raws from book table
*/
func (s *service) GetBooks() ([]*models.Book, error) {
	var books []*models.Book
	result := s.DB.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

/*
for given id, fetch only one entry of book in books table.
*/
func (s *service) GetOneBookById(id uint) (*models.Book, error) {
	var book *models.Book
	result := s.DB.Where("id = ?", id).Find(&book)

	if result.Error != nil || book.ID != id {
		return nil, errors.New("An error occured fetching book for given id")
	}

	return book, nil
}

/*
for given data, update book entry in books table.
*/
func (s *service) UpdateBook(book *models.Book) (*models.Book, error) {
	// update book
	result := s.DB.Save(book)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	log.Println("Book Updated", *book)
	return book, nil
}

/*
for given id, delete book data in books table.
*/
func (s *service) DeleteBook(id uint) (*models.Book, error) {
	book := &models.Book{}
	book.ID = id
	result := s.DB.Where("id = ?", id).Unscoped().Delete(book)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("error occured")
	}
	log.Println("book deleted", *book)
	return book, nil
}

// TODO: Implement Error Handling at service level
