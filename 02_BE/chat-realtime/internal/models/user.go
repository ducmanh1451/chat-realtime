package models

import (
	"time"
)

type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Email       string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password    string     `gorm:"type:varchar(20);not null" json:"password"`
	FullName    string     `gorm:"type:varchar(100)" json:"full_name"`
	Gender      *int       `gorm:"type:tinyint" json:"gender"`
	DateOfBirth *time.Time `gorm:"type:datetime" json:"date_of_birth"`
	CreatedAt   *time.Time `gorm:"autoCreateTime;type:datetime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime;type:datetime" json:"update_at"`
	DeletedAt   *time.Time `gorm:"index;type:datetime" json:"deleted_at"`
}
