package model

import "time"

type BaseMod struct {
	ID        uint64           	`gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time			`gorm:"autoCreateTime"`
	UpdatedAt time.Time			`gorm:"autoUpdateTime"`
}
