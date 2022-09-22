package pixiv_service

import "time"

type ArtworkDTO struct {
	Pid     int64  `json:"pid"`
	ArtId   int64  `json:"art_id"`
	ArtType string `json:"art_type"`

	Title         string   `json:"title"`
	Description   string   `json:"description"`
	PageCount     int64    `json:"page_count"`
	ViewCount     int64    `json:"view_count"`
	LikeCount     int64    `json:"like_count"`
	BookmarkCount int64    `json:"bookmark_count"`
	Tags          []string `json:"tags"`

	CreateDate *time.Time `json:"create_date"`
}
