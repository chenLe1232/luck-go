package models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"type:varchar(32);not null;unique"`
	Password  string `gorm:"type:varchar(128);not null"`
	Email     string `gorm:"type:varchar(128);unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
