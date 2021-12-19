package model

type PixivAccountMod struct {
	BaseMod

	Sid      string `gorm:"type:varchar(64);not null"`
	IsEnable bool   `gorm:"not null,default:false"`
}
