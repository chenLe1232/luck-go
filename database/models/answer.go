package models

import "time"

type Answer struct {
	ID        uint   `gorm:"primarykey"`
	Uid       string `gorm:"type:varchar(128);not null;unique"`
	CommitID  string `gorm:"type:varchar(128);not null;unique"` // 会话ID 一个用户可以有多个会话
	CreatedAt time.Time
	UpdatedAt time.Time
}
