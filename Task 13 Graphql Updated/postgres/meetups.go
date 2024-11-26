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

// Not Implemented.
func (m *MeetupRepo) UpdateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	return nil, nil
}

func (m *MeetupRepo) GetMeetupsBySearchText(searchText string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	// all descriptions starting with searchText
	err := m.DB.Model(&meetups).Where("description LIKE ?", searchText+"%").Select()

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
