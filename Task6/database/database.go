package database

import (
	"fmt"
	"task6MuxGorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/* Establish Connection With database */
func InitialiseDB(config *models.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v", config.HOST, config.DB_PORT, config.DB_USER, config.DB_USER_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
