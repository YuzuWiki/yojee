package apis

import (
	"encoding/json"
	"time"
)

type workDTO struct {
	ID             int64     `json:"id,string"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Description    string    `json:"description"`
	UserID         int64     `json:"userId,string"`
	UserName       string    `json:"userName"`
	Width          int       `json:"width"`
	Height         int       `json:"height"`
	PageCount      int       `json:"pageCount"`
	IsBookmarkable bool      `json:"isBookmarkable"`
	Alt            string    `json:"alt"`
	CreateDate     time.Time `json:"createDate"`
	UpdateDate     time.Time `json:"updateDate"`
}
type BookmarkDTO struct {
	Works []workDTO `json:"works"`
	Total int       `json:"total"`
}

/*
TagsDTO
	https://www.pixiv.net/ajax/user/32835219/illusts/bookmark/tags?lang=zh
*/
type bookMarkTagDTO struct {
	Tag string `json:"tag"`
	Cnt int32  `json:"cnt"`
}

type BookMarkTagsDTO struct {
	Private             []bookMarkTagDTO `json:"private"`
	Public              []bookMarkTagDTO `json:"public"`
	TooManyBookmark     bool             `json:"tooManyBookmark"`
	TooManyBookmarkTags bool             `json:"tooManyBookmarkTags"`
}

/*
ProfileAllDTO
	profile all
*/
type pickupDTO struct {
	Types       string   `json:"types"`
	Id          int64    `json:"id,string"`
	Tags        []string `json:"tags"`
	UserId      int64    `json:"userId,string"`
	UserName    string   `json:"userName"`
	Alt         string   `json:"alt"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	ContentUrl  string   `json:"contentUrl"`
}

type countDTO struct {
	Illust int `json:"illust"`
	Novel  int `json:"novel"`
}

type bookmarkCountDTO struct {
	Public  countDTO `json:"public"`
	Private countDTO `json:"private"`
}

type IllustMapDTO map[string]struct{}

func (dto *IllustMapDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]struct{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type MangaMapDTO map[string]struct{}

func (dto *MangaMapDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]struct{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type ProfileAllDTO struct {
	Illusts       IllustMapDTO     `json:"illusts"`
	Manga         MangaMapDTO      `json:"manga"`
	Pickup        []pickupDTO      `json:"pickup"`
	BookmarkCount bookmarkCountDTO `json:"bookmark_count"`
}

/*
ProfileTopDTO
	https://www.pixiv.net/ajax/user/7038833/profile/top?lang=zh
*/

type illustDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	IllustType      int       `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	PageCount       int       `json:"pageCount"`
	IsBookmarkable  bool      `json:"isBookmarkable"`
	Alt             string    `json:"alt"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	IsUnlisted      bool      `json:"isUnlisted"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type mangaDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	IllustType      int       `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	PageCount       int       `json:"pageCount"`
	IsBookmarkable  bool      `json:"isBookmarkable"`
	Alt             string    `json:"alt"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	IsUnlisted      bool      `json:"isUnlisted"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type novelDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	URL             string    `json:"url"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	ProfileImageURL string    `json:"profileImageUrl"`
	TextCount       int       `json:"textCount"`
	Description     string    `json:"description"`
	IsBookmarkable  bool      `json:"isBookmarkable"`
	BookmarkCount   int       `json:"bookmarkCount"`
	IsOriginal      bool      `json:"isOriginal"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	IsMasked        bool      `json:"isMasked"`
	SeriesID        string    `json:"seriesId"`
	SeriesTitle     string    `json:"seriesTitle"`
	IsUnlisted      bool      `json:"isUnlisted"`
}

type extraDataDTO struct {
	Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Canonical   string `json:"canonical"`
		Ogp         struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"ogp"`
		Twitter struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Card        string `json:"card"`
		} `json:"twitter"`
		DescriptionHeader string `json:"descriptionHeader"`
	} `json:"meta"`
}

type illustMapDTO map[string]illustDTO

func (dto *illustMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]illustDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type novelMapDTO map[string]novelDTO

func (dto *novelMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]novelDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type mangaMapDTO map[string]mangaDTO

func (dto *mangaMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]mangaDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type ProfileTopDTO struct {
	Illusts   illustMapDTO `json:"illusts"`
	Manga     mangaMapDTO  `json:"manga"`
	Novels    novelMapDTO  `json:"novels"`
	ExtraData extraDataDTO `json:"extra_data"`
}

/*
ExtraDTO
	https://www.pixiv.net/ajax/user/extra?lang=zh
*/
type ExtraDTO struct {
	Following    int32 `json:"following"`
	Followers    int32 `json:"followers"`
	MyPixivCount int32 `json:"mypixivCount"`
}

/*
	follow last
*/
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

type tagTranslationDTO map[string]tagDTO

func (dto *tagTranslationDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]tagDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type thumbnailDTO struct {
	Illust []illustDTO `json:"illust"`
	Novel  []novelDTO  `json:"novel"`
}

type FollowLatestDTO struct {
	Page           pageDTO           `json:"page"`
	TagTranslation tagTranslationDTO `json:"tagTranslation"`
	Thumbnails     thumbnailDTO      `json:"thumbnails"`
}

/*
	following
*/
type followUserDTO struct {
	UserID          int64       `json:"userId,string"`
	UserName        string      `json:"userName"`
	ProfileImageURL string      `json:"profileImageUrl"`
	UserComment     string      `json:"userComment"`
	Following       bool        `json:"following"`
	Followed        bool        `json:"followed"`
	IsBlocking      bool        `json:"isBlocking"`
	IsMypixiv       bool        `json:"isMypixiv"`
	Illusts         []illustDTO `json:"illusts"`
	AcceptRequest   bool        `json:"acceptRequest"`
}

type FollowDTO struct {
	Total int             `json:"total"`
	Users []followUserDTO `json:"users"`
}
