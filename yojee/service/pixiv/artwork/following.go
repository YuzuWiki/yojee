package artwork

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

const (
	StatusShow = "show"
	StatusHide = "hide"
)

type FollowingAPI struct{}

type IllustDTO struct {
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	IllustType      int         `json:"illustType"`
	URL             string      `json:"url"`
	Description     string      `json:"description"`
	Tags            []string    `json:"tags"`
	UserID          string      `json:"userId"`
	UserName        string      `json:"userName"`
	Width           int         `json:"width"`
	Height          int         `json:"height"`
	PageCount       int         `json:"pageCount"`
	BookmarkData    interface{} `json:"bookmarkData"`
	Alt             string      `json:"alt"`
	CreateDate      time.Time   `json:"createDate"`
	UpdateDate      time.Time   `json:"updateDate"`
	IsUnlisted      bool        `json:"isUnlisted"`
	IsMasked        bool        `json:"isMasked"`
	ProfileImageURL string      `json:"profileImageUrl"`
}

type UserDTO struct {
	UserID          string      `json:"userId"`
	UserName        string      `json:"userName"`
	ProfileImageURL string      `json:"profileImageUrl"`
	UserComment     string      `json:"userComment"`
	Following       bool        `json:"following"`
	Followed        bool        `json:"followed"`
	IsBlocking      bool        `json:"isBlocking"`
	IsMypixiv       bool        `json:"isMypixiv"`
	Illusts         []IllustDTO `json:"illusts"`
	AcceptRequest   bool        `json:"acceptRequest"`
}

type FollowingDTO struct {
	Total int       `json:"total"`
	Users []UserDTO `json:"users"`
}

func (f FollowingAPI) path(uid int32) string {
	return fmt.Sprintf("ajax/user/%d/following", uid)
}

func (f FollowingAPI) find(ctx context.Context, uid int32, rest string, limit int, offset int) (flowing *FollowingDTO, err error) {
	c := client.For(ctx)

	// 添加参数
	query := netUrl.Values{}
	query.Set("offset", strconv.Itoa(offset))
	query.Set("limit", strconv.Itoa(limit))
	query.Set("rest", rest)
	query.Set("tag", "")
	query.Set("lang", "zh")

	queryUrl := c.EndpointULR(f.path(uid), &query).String()
	resp, err := c.GetWithContext(ctx, queryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := struct {
		Error bool         `json:"error"`
		Body  FollowingDTO `json:"body"`
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Print(string(body))
		return nil, err
	}

	if data.Error == true {
		return nil, errors.New(fmt.Sprintf("API Response error: url = %s;  data = %s", queryUrl, body))
	}

	return &data.Body, nil
}

func (f FollowingAPI) FindShow(ctx context.Context, uid int32, limit int, offset int) (flowing *FollowingDTO, err error)  {
	return f.find(ctx, uid, StatusShow, limit, offset)
}

func (f FollowingAPI) FindHide(ctx context.Context, uid int32, limit int, offset int) (flowing *FollowingDTO, err error)  {
	return f.find(ctx, uid, StatusHide, limit, offset)
}



