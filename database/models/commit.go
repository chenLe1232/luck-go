package models

import "time"

// 创建一个会话记录 一个会话记录包含多个消息，每个消息内容是 question answer， 创建时间 更新时间
type Commit struct {
	ID        uint   `gorm:"primarykey"`
	CommitID  string `gorm:"type:varchar(128);not null;unique"`
	Question  string `gorm:"type:text;not null"`
	Answer    string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
