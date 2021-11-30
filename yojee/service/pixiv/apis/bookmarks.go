package apis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/like9th/yojee/yojee/service/pixiv/client"
	"io/ioutil"
	netUrl "net/url"
	"strconv"
	"time"
)

type BookmarkAPI struct {}

type WorkDTO struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Sl int `json:"sl"`
	URL string `json:"url"`
	Description string `json:"description"`
	UserID string `json:"userId"`
	UserName string `json:"userName"`
	Width int `json:"width"`
	Height int `json:"height"`
	PageCount int `json:"pageCount"`
	IsBookmarkable bool `json:"isBookmarkable"`
	Alt string `json:"alt"`
	CreateDate time.Time `json:"createDate"`
	UpdateDate time.Time `json:"updateDate"`
}
type BookmarkDTO struct {
	Works []WorkDTO `json:"works"`
	Total int `json:"total"`
}

func (api BookmarkAPI)path(uid int32) string  {
	return fmt.Sprintf("ajax/user/%d/illusts/bookmarks", uid)
}

func (api BookmarkAPI) find(ctx context.Context, uid int32, tag string, offset int, limit int, rest string) (*BookmarkDTO, error){
	c  := client.For(ctx)

	query := netUrl.Values{}
	query.Set("tag", tag)
	query.Set("limit", strconv.Itoa(limit))
	query.Set("offset", strconv.Itoa(offset))
	query.Set("rest", rest)
	query.Set("lang", "zh")

	queryUrl := c.EndpointULR(api.path(uid), &query).String()
	resp, err := c.GetWithContext(ctx, queryUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := struct {
		Error bool         `json:"error"`
		Body  BookmarkDTO  `json:"body"`
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if data.Error == true {
		return nil, errors.New(fmt.Sprintf("API Response error: url = %s;  data = %s", queryUrl, body))
	}
	return &data.Body, nil
}
func (api BookmarkAPI) FindShow(ctx context.Context, uid int32, tag string, offset int, limit int) (*BookmarkDTO, error){
	return api.find(ctx, uid, tag, offset, limit, StatusShow)
}

func (api BookmarkAPI) FindHide(ctx context.Context, uid int32, tag string, offset int, limit int) (*BookmarkDTO, error){
	return api.find(ctx, uid, tag, offset, limit, StatusHide)
}