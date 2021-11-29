package artwork

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/like9th/yojee/yojee/service/pixiv/client"
)

type UserAPI struct{}

type UserProfileDTO struct {
	Illusts     map[string]interface{} `json:"illusts"`
	Manga       []interface{}          `json:"manga"`
	Novels      []interface{}          `json:"novels"`
	MangaSeries []interface{}          `json:"mangaSeries"`
	NovelSeries []interface{}          `json:"novelSeries"`
	Pickup      []struct {
		Type       string `json:"type"`
		UserName   string `json:"userName"`
		ContentURL string `json:"contentUrl"`
	} `json:"pickup"`
	BookmarkCount struct {
		Public struct {
			Illust int `json:"illust"`
			Novel  int `json:"novel"`
		} `json:"public"`
		Private struct {
			Illust int `json:"illust"`
			Novel  int `json:"novel"`
		} `json:"private"`
	} `json:"bookmarkCount"`
	ExternalSiteWorksStatus struct {
		Booth    bool `json:"booth"`
		Sketch   bool `json:"sketch"`
		VroidHub bool `json:"vroidHub"`
	} `json:"externalSiteWorksStatus"`
	Request struct {
		ShowRequestTab     bool `json:"showRequestTab"`
		ShowRequestSentTab bool `json:"showRequestSentTab"`
		PostWorks          struct {
			Artworks []interface{} `json:"artworks"`
			Novels   []interface{} `json:"novels"`
		} `json:"postWorks"`
	} `json:"request"`
}

func (u UserAPI) path(userID int32) string {
	return fmt.Sprintf("ajax/user/%d/profile/all", userID)
}

func (u UserAPI) UserProfile(ctx context.Context, userID int32) (*UserProfileDTO, error) {
	c := client.For(ctx)

	queryUrl := c.EndpointULR(u.path(userID), nil).String()
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
		Error bool           `json:"error"`
		Body  UserProfileDTO `json:"body"`
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.New(fmt.Sprintf("API Response error: url = %s;  data = %s", queryUrl, body))
	}

	if data.Error == true {
		return nil, errors.New(fmt.Sprintf("API Response error: url = %s;  data = %s", queryUrl, body))
	}

	return &data.Body, nil
}
