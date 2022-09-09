package apis

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/dtos"
)

func GetAccountInfo(ctx pixiv_v2.IContext, uid int64) (_ *dtos.UserInfoDTO, err error) {
	var (
		query *pixiv_v2.Query
		c     pixiv_v2.IClient
	)
	if query, err = pixiv_v2.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Body(c.Get, pixiv_v2.Path("/users", uid), query, nil)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	user := &struct {
		User map[string]dtos.UserInfoDTO `json:"user"`
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

func GetProfileAll(ctx pixiv_v2.IContext, uid int64) (body *dtos.AllProfileDTO, err error) {
	var c pixiv_v2.IClient
	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax/user/", uid, "/profile", All), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}

	return body, nil
}

func GetProfileTop(ctx pixiv_v2.IContext, uid int64) (body *dtos.TopProfileDTO, err error) {
	var c pixiv_v2.IClient
	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax/user/", uid, "/profile", Top), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}

	return body, nil
}
