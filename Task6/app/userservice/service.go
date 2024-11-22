package userservice

/*
	User Services
	- create a new User
	- get all books
	- get one User by id
	- update one User
	- delete one User
*/

import (
	"errors"
	"fmt"
	"log"
	"task6MuxGorm/app"
	"task6MuxGorm/models"

	"gorm.io/gorm"
)

/*
define all services a bookservice needs.
*/
type Service interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetOneUserById(id uint64) (*models.User, error)
	UpdateUser(User *models.User) (*models.User, error)
	DeleteUser(id uint64) (*models.User, error)
	IssueBook(userId uint64, bookId uint64) (*models.User, error)
}

/*
instance of db with required data which needed for performing its services.
here only database connection required, as it will be doing only raw operations.
*/
type service struct { // Implement Service Interface
	DB *gorm.DB
}

/*
create new User service
get DB from app and initialise it.
*/
func New(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// TODO: Implement Service Interface

/*
for given data, insert raw in User table.
*/
func (s *service) CreateUser(User *models.User) (*models.User, error) {
	result := s.DB.Create(User)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	log.Println("User inserted", *User)
	return User, nil
}

/*
Get all the raws from User table
*/
func (s *service) GetUsers() ([]*models.User, error) {
	var users []*models.User
	result := s.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}


	s.DB.Preload("Books").Find(&users)
	return users, nil
}

/*
for given id, fetch only one entry of User in books table.
*/
func (s *service) GetOneUserById(id uint64) (*models.User, error) {
	var user *models.User
	result := s.DB.Where("id = ?", id).Find(&user)

	if result.Error != nil || user.ID != id {
		return nil, errors.New("An error occured fetching User for given id")
	}

	return user, nil
}

/*
for given data, update User entry in books table.
*/
func (s *service) UpdateUser(user *models.User) (*models.User, error) {
	// update User
	result := s.DB.Save(user)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	log.Println("User Updated", *user)
	return user, nil
}

/*
for given id, delete User data in books table.
*/
func (s *service) DeleteUser(id uint64) (*models.User, error) {
	user := &models.User{}
	user.ID = id
	result := s.DB.Unscoped().Delete(user)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("error occured")
	}
	log.Println("User deleted", *user)
	return user, nil
}

// update User struct with latest data and return it with issued book.
func (s *service) IssueBook(userId uint64, bookId uint64) (*models.User, error) {
	// issue book
	// apply row level lock

	// Apply row-level lock by using `Select` with `For` clause
	// Get data from row after locking the row of book

	// TODO: Check if user already has issued book

	// ASSUMING no one can update association table
	var bookUser models.BookUser
	s.DB.Raw("SELECT * FROM task6muxgorm.usersbooks WHERE user_id = ? AND book_id = ?;", userId, bookId).Scan(&bookUser)

	fmt.Println(bookUser)
	if bookUser.BookId != 0 {
		return nil, errors.New("User already issued this book")
	}

	// Begin Transaction For Issuing Book // Lock Given Book And Given User For Further Update
	tx := s.DB.Begin()

	if tx.Error != nil {
		return nil, errors.New("failed to begin transaction")
	}

	var user models.User
	var book models.Book

	// Lock User Row
	query := "SELECT * FROM task6muxgorm.users WHERE id = ? FOR UPDATE;"
	result := tx.Raw(query, userId).Scan(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, errors.New("Can not lock user row")
	}

	// Lock Book Row
	query = "SELECT * FROM task6muxgorm.books WHERE id = ? FOR UPDATE;"
	result = tx.Raw(query, bookId).Scan(&book)
	if result.Error != nil || book.ID == 0 {
		tx.Rollback()
		return nil, errors.New("Can not lock book row or book does not exists")
	}

	// ? Do I need to lock association table of book and user here. No ig i am not trying to update it anyway. I will insert data at most.

	// If Book Quantity Greater Then 0 then issue it
	if book.Quantity > 0 {
		// update book quantity to less 1
		book.Quantity = book.Quantity - 1
		bookSaveResult := tx.Save(book)
		if bookSaveResult.Error != nil {
			tx.Rollback()
			return nil, errors.New("not able to update quantity of book")
		}
		// update user issued book to current book
		query := "INSERT INTO task6muxgorm.usersbooks(user_id, book_id) VALUES(?,?);"
		tx.Exec(query, userId, bookId)
	} else {
		// can't issue book
		return nil, errors.New("Can't issue book because quantity is less then 0")
	}

	// Release Locks
	tx.Commit()

	s.DB.Preload("Books").Find(&user)
	return &user, nil
}

// TODO: Implement Error Handling at service level
