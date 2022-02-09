package model

import "time"

type PixivUserMod struct {
	BaseMod

	UUID     string `gorm:"type:varchar(128)" json:"uuid"`
	PID      int64  `gorm:"type:bigint;column:pid" json:"pid"`
	Name     string `gorm:"type:varchar(256);column:name" json:"name"`
	NickName string `gorm:"type:varchar(256);column:nick_name" json:"nick_name"`

	Avatar    string `gorm:"type:varchar(256);column:avatar" json:"avatar"`
	Following int    `gorm:"type:int;default:0;column:following;comment:'关注数量'" json:"following"`
	Followers int    `gorm:"type:int;default:0;column:followers;'粉丝数量'" json:"followers"`
}

type PixivFollowMod struct {
	BaseMod

	PID         int64      `gorm:"type:bigint;column:pid" json:"pid"`
	FollowedPid int64      `gorm:"type:bigint;column:followed_pid" json:"followed_pid"`
	FollowedAt  *time.Time `gorm:"type:timestamp;column:followed_at"  json:"followed_at"`
}
