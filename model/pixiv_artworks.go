package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

type PixivArtworkMod struct {
	ID        uint64     `gorm:"type:timestamp;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

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

func (PixivArtworkMod) Find(artType string, pid int64) (artworks *[]PixivArtworkTagMod, err error) {
	if err = global.DB().Exec("SELECT * FROM pixiv_artwork WHERE pid=? AND art_type=? AND is_deleted=false;", pid, artType).Find(artworks).Error; err != nil {
		return nil, err
	}
	return
}

func (PixivArtworkMod) Insert(data dtos.ArtworkDTO) (int64, error) {
	row := PixivArtworkMod{
		Pid:           data.Pid,
		ArtId:         data.ArtId,
		ArtType:       data.ArtType,
		Title:         data.Title,
		Description:   data.Description,
		ViewCount:     data.ViewCount,
		LikeCount:     data.LikeCount,
		BookmarkCount: data.BookmarkCount,
		CreateDate:    &data.CreateDate,
	}
	if err := global.DB().FirstOrCreate(&row, PixivArtworkMod{Pid: data.Pid, ArtType: data.ArtType, ArtId: data.ArtId}).Error; err != nil {
		global.Logger.Error().Msg(fmt.Sprintf("insert illust(%d) error,  %s", data.ArtId, err.Error()))
		return 0, err
	}
	return int64(row.ID), nil
}
