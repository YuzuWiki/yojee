package model

import "time"

type BaseMod struct {
	ID        uint64     `gorm:"primaryKey,autoIncrement"`
	CreatedAt *time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli"`
	IsDeleted bool       `gorm:"type:bool;default:false"`
}
