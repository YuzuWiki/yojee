package model

import (
	"strings"
	"time"

	"github.com/YuzuWiki/yojee/global"
)

type PixivAccountMod struct {
	ID        uint64     `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	Pid        int64  `gorm:"type:bigint;column:pid" json:"pid"`
	Name       string `gorm:"type:varchar(256);column:name" json:"name"`
	Avatar     string `gorm:"type:varchar(512);column:avatar" json:"avatar"`
	Region     string `gorm:"type:varchar(128);column:region" json:"region"`
	Gender     string `gorm:"type:varchar(128);column:gender" json:"gender"`
	BirthDay   string `gorm:"type:varchar(128);column:birth_day" json:"birthDay"`
	Job        string `gorm:"type:varchar(128);column:job" json:"job"`
	Following  int32  `gorm:"type:int;default:0;column:following;comment:'关注数量'" json:"following"`
	FanboxUrl  string `gorm:"type:varchar(256);column:fanbox_url" json:"fanbox_url"`
	ArtUpdated int64  `gorm:"type:bigint;default:0;column:art_updated;comment:'作品最后一次更细那时间'" json:"art_updated"`
}

func (PixivAccountMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_account"}, ".")
}

func (PixivAccountMod) Insert(pid int64, name string, avatar string, region string, gender string, following int32) (int64, error) {
	row := &PixivAccountMod{
		Pid:       pid,
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
	ID        uint64     `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	PID         int64 `gorm:"type:bigint;column:pid" json:"pid"`
	FollowedPid int64 `gorm:"type:bigint;column:followed_pid" json:"followed_pid"`
}

func (PixivFollowMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_follow"}, ".")
}

func (PixivFollowMod) MarkFollowing(pid int64, followedPid int64) (int64, error) {
	row := &PixivFollowMod{
		PID:         pid,
		FollowedPid: followedPid,
	}
	if err := global.DB().Create(row).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}
