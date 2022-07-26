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
