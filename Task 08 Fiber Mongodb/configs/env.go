package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}

func ENV_MONGO_URI() string {
	return os.Getenv("MONGO_URI")
}

func ENV_PORT() string {
	return os.Getenv("PORT")
}

func ENV_DATBASE() string {
	return os.Getenv("DATABASE")
}
