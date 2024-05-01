package postgres

import (
	"meetmeup/models"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (m *UserRepo) GetUserById(id string) (*models.User, error) {
	var user models.User

	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserRepo) GetMeetupsByUserId(id string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Where("user_id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
