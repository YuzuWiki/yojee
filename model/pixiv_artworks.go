package model

import (
	"strings"
	"time"

	"github.com/YuzuWiki/yojee/global"
)

type PixivArtworkMod struct {
	ID        int64 `gorm:"type:bigint;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt int64 `gorm:"type:bigint;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"type:bigint;autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	IsDeleted bool  `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	Pid     int64  `gorm:"type:bigint;column:pid" json:"pid"`
	ArtId   int64  `gorm:"type:bigint;column:art_id" json:"art_id"`
	ArtType string `gorm:"type:varchar(64);column:art_type" json:"art_type"`

	Title         string     `gorm:"type:text;column:title" json:"title"`
	Description   string     `gorm:"type:text;column:description" json:"description"`
	PageCount     int64      `gorm:"type:bigint;default:0;column:page_count" json:"page_count"`
	ViewCount     int64      `gorm:"type:bigint;default:0;column:view_count" json:"view_count"`
	LikeCount     int64      `gorm:"type:bigint;default:0;column:like_count" json:"like_count"`
	BookmarkCount int64      `gorm:"type:bigint;default:0;column:bookmark_count" json:"bookmark_count"`
	CreateDate    *time.Time `gorm:"type:timestamp;column:create_date" json:"create_date"`
}

func (PixivArtworkMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_artwork"}, ".")
}
