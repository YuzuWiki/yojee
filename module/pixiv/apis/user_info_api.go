package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	pixiv2 "github.com/YuzuWiki/yojee/module/pixiv"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type InfoAPI struct{}

func (api InfoAPI) Extra(ctx pixiv2.Context) (*ExtraDTO, error) {
	data, err := pixiv2.Request(ctx, http.MethodGet, pixiv2.Path("/ajax/user", "extra"), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ExtraDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (api InfoAPI) Info(ctx pixiv2.Context, pid int64) (*UserInfoDTO, error) {
	query, err := pixiv2.NewQuery(map[string]interface{}{"lang": "en"})
	if err != nil {
		return nil, err
	}

	data, err := pixiv2.Request(ctx, http.MethodGet, pixiv2.Path("/users", pid), query, nil)
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
