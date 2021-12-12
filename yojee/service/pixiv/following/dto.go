package following

import "time"

type pageDTO struct {
	Ids  []int32       `json:"ids"`
	Tags []interface{} `json:"tags"`
}

type tagDTO struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	ZhTw   string `json:"zh_tw"`
	Romaji string `json:"romaji"`
}

type illustDTO struct {
	ID                      string                 `json:"id"`
	Title                   string                 `json:"title"`
	IllustType              int                    `json:"illustType"`
	XRestrict               int                    `json:"xRestrict"`
	Restrict                int                    `json:"restrict"`
	Sl                      int                    `json:"sl"`
	URL                     string                 `json:"url"`
	Description             string                 `json:"description"`
	Tags                    []string               `json:"tags"`
	UserID                  string                 `json:"userId"`
	UserName                string                 `json:"userName"`
	Width                   int                    `json:"width"`
	Height                  int                    `json:"height"`
	PageCount               int                    `json:"pageCount"`
	IsBookmarkable          bool                   `json:"isBookmarkable"`
	Alt                     string                 `json:"alt"`
	TitleCaptionTranslation map[string]interface{} `json:"titleCaptionTranslation"`
	CreateDate              time.Time              `json:"createDate"`
	UpdateDate              time.Time              `json:"updateDate"`
	IsUnlisted              bool                   `json:"isUnlisted"`
	IsMasked                bool                   `json:"isMasked"`
	Urls                    map[string]string      `json:"urls"`
	ProfileImageURL         string                 `json:"profileImageUrl"`
}

type novelDTO struct {
	ID                      string                 `json:"id"`
	Title                   string                 `json:"title"`
	XRestrict               int                    `json:"xRestrict"`
	Restrict                int                    `json:"restrict"`
	URL                     string                 `json:"url"`
	Tags                    []string               `json:"tags"`
	UserID                  string                 `json:"userId"`
	UserName                string                 `json:"userName"`
	ProfileImageURL         string                 `json:"profileImageUrl"`
	TextCount               int                    `json:"textCount"`
	Description             string                 `json:"description"`
	IsBookmarkable          bool                   `json:"isBookmarkable"`
	BookmarkCount           int                    `json:"bookmarkCount"`
	IsOriginal              bool                   `json:"isOriginal"`
	TitleCaptionTranslation map[string]interface{} `json:"titleCaptionTranslation"`
	CreateDate              time.Time              `json:"createDate"`
	UpdateDate              time.Time              `json:"updateDate"`
	IsMasked                bool                   `json:"isMasked"`
	IsUnlisted              bool                   `json:"isUnlisted"`
}

type thumbnailDTO struct {
	Illust []illustDTO `json:"illust"`
	Novel  []novelDTO  `json:"novel"`
}

type FollowLatestDTO struct {
	Page           pageDTO           `json:"page"`
	TagTranslation map[string]tagDTO `json:"tag_translation"`
	Thumbnails     thumbnailDTO      `json:"thumbnails"`
}
