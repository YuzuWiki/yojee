package model

import "time"

type PixivUserMod struct {
	BaseMod

	UUID     string `gorm:"type:varchar(128)" json:"uuid"`
	PID      int64  `gorm:"type:bigint;column:pid;comment:'pixiv id, 高级会员可变更'" json:"pid"`
	Name     string `gorm:"type:varchar(128);column:name" json:"name"`
	NickName string `gorm:"type:varchar(128);column:nick_name" json:"nick_name"`

	Avatar      string     `gorm:"type:varchar(256);column:avatar" json:"avatar"`
	Following   int        `gorm:"type:int;default:0;column:following;comment:'关注数量'" json:"following"`
	Followers   int        `gorm:"type:int;default:0;column:followers;'粉丝数量'" json:"followers"`
	FollowingAt *time.Time `gorm:"type:int;default:0;column:following_at;'关注时间'" json:"following_at"`
}
