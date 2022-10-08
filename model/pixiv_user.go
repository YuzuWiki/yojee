package model

import (
	"strings"

	"github.com/YuzuWiki/yojee/global"
)

type PixivAccountMod struct {
	ID        int64 `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt int64 `gorm:"type:bigint;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"type:bigint;autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	IsDeleted bool  `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

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

type PixivFollowMod struct {
	ID        int64 `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt int64 `gorm:"type:bigint;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"type:bigint;autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	IsDeleted bool  `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	PID         int64 `gorm:"type:bigint;column:pid" json:"pid"`
	FollowedPid int64 `gorm:"type:bigint;column:followed_pid" json:"followed_pid"`
}

func (PixivFollowMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_follow"}, ".")
}
