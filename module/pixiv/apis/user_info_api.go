package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

type InfoAPI struct{}

func (InfoAPI) Information(ctx pixiv.Context, pid int64) (*UserInfoDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{"lang": "en"})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/users", pid), query, nil)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfoDTO{}
	document.Find("#meta-preload-data").Each(func(i int, selection *goquery.Selection) {
		if err := json.Unmarshal([]byte(selection.Text()), userInfo); err != nil {
			return
		}
	})

	if userInfo.UserID == 0 {
		return nil, errors.New("not Found UserInfo")
	}
	return userInfo, nil
}

func (InfoAPI) Artwork(ctx pixiv.Context, pid int64) (*ProfileAllDTO, error) {
	return profileInfo(ctx, pid)
}
