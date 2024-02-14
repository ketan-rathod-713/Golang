// define and create service with interface
package bookservice

import (
	"log"
	"task6MuxGorm/app"
	"task6MuxGorm/models"

	"gorm.io/gorm"
)

type Service interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBooks() ([]*models.Book, error)
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(book *models.Book) (*models.Book, error)
}

// This service only requires DB connection for doing its job
type service struct { // Implement Service Interface
	DB *gorm.DB
}

// Create new service
func New(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// TODO: Implement Service Interface
func (s *service) CreateBook(book *models.Book) (*models.Book, error) {
	s.DB.Create(book)
	log.Println("Book inserted", *book)
	return book, nil
}

func (s *service) GetBooks() ([]*models.Book, error) {
	var books []*models.Book
	s.DB.Find(&books)
	return books, nil
}

func (s *service) UpdateBook(book *models.Book) (*models.Book, error) {
	// update book
	s.DB.Save(book)
	log.Println("Book Updated", *book)
	return book, nil
}

func (s *service) DeleteBook(book *models.Book) (*models.Book, error) {

	log.Println("book deleted", *book)
	return book, nil
}
