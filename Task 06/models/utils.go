package models

import (
	"fmt"
)

type Config struct {
	PORT             string
	DB_PORT          string
	DATABASE         string
	HOST             string
	DB_USER          string
	DB_USER_PASSWORD string
	DB_SCHEMA_NAME   string
}

// Error schema to show for userA
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Status Code:%v, Message: %v", e.Code, e.Message)
}
