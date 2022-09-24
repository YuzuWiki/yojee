package pixiv_service

import (
	"encoding/json"
	"strings"
	"time"
)

type ArtworkDTO struct {
	Pid     int64  `gorm:"column:pid" json:"pid"`
	ArtId   int64  `gorm:"column:art_id" json:"art_id"`
	ArtType string `gorm:"column:art_type" json:"art_type"`

	Title         string        `gorm:"column:title" json:"title"`
	Description   string        `gorm:"column:description" json:"description"`
	PageCount     int64         `gorm:"column:page_count" json:"page_count"`
	ViewCount     int64         `gorm:"column:view_count" json:"view_count"`
	LikeCount     int64         `gorm:"column:like_count" json:"like_count"`
	BookmarkCount int64         `gorm:"column:bookmark_count" json:"bookmark_count"`
	CreateDate    *time.Time    `gorm:"column:create_date" json:"create_date"`
	Tags          artworkTagDTO `gorm:"column:tags" json:"tags"`
}

type artworkTagDTO []uint8

func (dto *artworkTagDTO) MarshalJSON() ([]byte, error) {
	tagStr := string(*dto)
	if len(tagStr) == 0 {
		return []byte{'[', ']'}, nil
	}
	return json.Marshal(strings.Split(tagStr, ","))
}
