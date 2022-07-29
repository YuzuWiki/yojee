package user

import "time"

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
