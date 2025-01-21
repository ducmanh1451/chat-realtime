package models

import (
	"time"
)

type User struct {
	ID          int        `json:"id"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	FullName    string     `json:"full_name"`
	Gender      *int       `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
