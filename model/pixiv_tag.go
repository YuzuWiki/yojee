package model

import (
	"strings"
	"time"

	"github.com/YuzuWiki/yojee/global"
)

const _TagIdKey = "pixiv:tag:id:"

type PixivTagMod struct {
	ID        uint64     `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	TagId  int64  `gorm:"type:bigint;column:tag_id" json:"tag_id"`
	Jp     string `gorm:"type:VARCHAR(512);column:jp" json:"jp"`
	En     string `gorm:"type:VARCHAR(512);column:en" json:"en"`
	Ko     string `gorm:"type:VARCHAR(512);column:ko" json:"ko"`
	Zh     string `gorm:"type:VARCHAR(512);column:zh" json:"zh"`
	Romaji string `gorm:"type:VARCHAR(512);column:romaji" json:"romaji"`
}

func (PixivTagMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_tag"}, ".")
}

type PixivTagTreeMod struct {
	ID        uint64     `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	ParentId int64  `gorm:"type:bigint;column:parent_id" json:"parent_id"`
	ParentJp string `gorm:"type:VARCHAR(512);column:parent_jp" json:"parent_jp"`

	TagId int64  `gorm:"type:bigint;column:tag_id" json:"tag_id"`
	TagJp string `gorm:"type:VARCHAR(512);column:tag_jp" json:"tag_jp"`
}

func (PixivTagTreeMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_tag_tree"}, ".")
}

type PixivArtworkTagMod struct {
	ID        uint64     `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	Pid     int64  `gorm:"type:bigint;column:pid" json:"pid"`
	ArtId   int64  `gorm:"type:bigint;column:art_id" json:"art_id"`
	ArtType string `gorm:"type:varchar(64);column:art_type" json:"art_type"`

	TagId int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivArtworkTagMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_artwork_tag"}, ".")
}
