package model

import "time"

type PixivIllustMod struct {
	BaseMod

	Pid int64 `gorm:"type:bigint;column:pid" json:"pid"`

	IllustId      int64      `gorm:"type:bigint;column:illust_id" json:"illust_id"`
	Title         string     `gorm:"type:text;column:title" json:"title"`
	Description   string     `gorm:"type:text;column:description" json:"description"`
	ViewCount     int64      `gorm:"type:bigint;default:0;column:view_count" json:"view_count"`
	LikeCount     int64      `gorm:"type:bigint;default:0;column:like_count" json:"like_count"`
	BookmarkCount int64      `gorm:"type:bigint;default:0;column:bookmark_count" json:"bookmark_count"`
	CreateDate    *time.Time `gorm:"type:timestamp;column:create_date" json:"create_date"`
}

func (PixivIllustMod) TableName() string {
	return "pixiv_illust"
}

type PixivMangaMod struct {
	BaseMod

	PID int64 `gorm:"type:bigint;column:pid" json:"pid"`

	MangaId       int64      `gorm:"type:bigint;column:manga_id" json:"manga_id"`
	Title         string     `gorm:"type:text;column:title" json:"title"`
	Description   string     `gorm:"type:text;column:description" json:"description"`
	PageCount     int64      `gorm:"type:bigint;default:0;column:page_count" json:"page_count"`
	ViewCount     int64      `gorm:"type:bigint;default:0;column:view_count" json:"view_count"`
	LikeCount     int64      `gorm:"type:bigint;default:0;column:like_count" json:"like_count"`
	BookmarkCount int64      `gorm:"type:bigint;default:0;column:bookmark_count" json:"bookmark_count"`
	CreateDate    *time.Time `gorm:"type:timestamp;column:create_date" json:"create_date"`
}

func (PixivMangaMod) TableName() string {
	return "pixiv_manga"
}

type PixivNovelMod struct {
	BaseMod

	PID int64 `gorm:"type:bigint;column:pid" json:"pid"`

	NovelId        int64      `gorm:"type:bigint;column:novel_id" json:"novel_id"`
	Title          string     `gorm:"type:text;column:title" json:"title"`
	Description    string     `gorm:"type:text;column:description" json:"description"`
	ChapterCount   int        `gorm:"type:int;default:0;column:chapter_count" json:"chapter_count"`
	WordageCount   int        `gorm:"type:int;default:0;column:wordage_count" json:"wordage_count"`
	CreateDate     *time.Time `gorm:"type:timestamp;column:create_date" json:"create_date"`
	LastUpdateDate *time.Time `gorm:"type:timestamp;column:last_update_date" json:"last_update_date"`
}

func (PixivNovelMod) TableName() string {
	return "pixiv_novel"
}
