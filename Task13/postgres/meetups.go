package postgres

import (
	"meetmeup/models"

	"github.com/go-pg/pg/v10"
)

type MeetupRepo struct {
	DB *pg.DB
}

func (m *MeetupRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {

	_, err := m.DB.Model(meetup).Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return meetup, nil
}
