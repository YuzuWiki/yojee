package model

import "time"

type BaseMod struct {
	ID        uint64     `gorm:"type:timestamp;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_delete" json:"is_deleted"`
}
