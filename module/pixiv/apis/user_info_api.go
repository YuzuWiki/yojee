package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
)

type InfoAPI struct{}

func (InfoAPI) Information(ctx pixiv.Context, pid int64) (*UserInfoDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{"lang": "jp"})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Body(ctx, http.MethodGet, pixiv.Path("/users", pid), query, nil)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	user := &struct {
		User map[string]UserInfoDTO `json:"user"`
	}{}
	document.Find("#meta-preload-data").Each(func(i int, selection *goquery.Selection) {
		var text string
		if _text, isExist := selection.Attr("content"); !isExist {
			return
		} else {
			text = _text
		}

		if err := json.Unmarshal([]byte(text), user); err != nil {
			global.Logger.Error().Err(err)
			return
		}
	})

	for _, info := range user.User {
		if info.UserID > 0 {
			return &info, nil
		}
	}
	return nil, fmt.Errorf("not Found UserInfo")
}

func (InfoAPI) Artwork(ctx pixiv.Context, pid int64) (*ProfileAllDTO, error) {
	return profileInfo(ctx, pid)
}
