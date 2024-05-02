package postgres

import (
	"meetmeup/models"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

// get user by id
func (m *UserRepo) GetUserById(id string) (*models.User, error) {
	var user models.User

	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Create a user
func (m *UserRepo) CreateUser(user models.User) (*models.User, error) {

	_,err := m.DB.Model(&user).Insert()
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// get all the users
func (m *UserRepo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := m.DB.Model(&users).Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// get meetups by user id
func (m *UserRepo) GetMeetupsByUserId(id string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Where("user_id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
