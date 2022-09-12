package model

import (
	"time"

	"github.com/YuzuWiki/yojee/global"
)

type PixivUserMod struct {
	BaseMod

	PID    int64  `gorm:"type:bigint;column:pid" json:"pid"`
	Name   string `gorm:"type:varchar(256);column:name" json:"name"`
	Avatar string `gorm:"type:varchar(512);column:avatar" json:"avatar"`
	Region string `gorm:"type:varchar(16);column:region" json:"region"`
	Gender string `gorm:"type:varchar(16);column:gender" json:"gender"`

	Following int32 `gorm:"type:int;default:0;column:following;comment:'关注数量'" json:"following"`
}

func (PixivUserMod) TableName() string {
	return "pixiv_user"
}

func (PixivUserMod) Insert(pid int64, name string, avatar string, region string, gender string, following int32) (int64, error) {
	row := &PixivUserMod{
		PID:       pid,
		Name:      name,
		Avatar:    avatar,
		Region:    region,
		Gender:    gender,
		Following: following,
	}
	if err := global.DB().Create(row).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}

type PixivFollowMod struct {
	BaseMod

	PID         int64      `gorm:"type:bigint;column:pid" json:"pid"`
	FollowedPid int64      `gorm:"type:bigint;column:followed_pid" json:"followed_pid"`
	FollowedAt  *time.Time `gorm:"type:timestamp;column:followed_at"  json:"followed_at"`
}

func (PixivFollowMod) TableName() string {
	return "pixiv_follow"
}

func (PixivFollowMod) MarkFollowing(pid int64, followedPid int64, followedAt time.Time) (int64, error) {
	row := &PixivFollowMod{
		PID:         pid,
		FollowedPid: followedPid,
		FollowedAt:  &followedAt,
	}
	if err := global.DB().Create(row).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}
