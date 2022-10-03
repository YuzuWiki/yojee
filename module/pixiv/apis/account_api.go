package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func GetAccountInfo(ctx pixiv.IContext, pid int64) (body *dtos.UserInfoDTO, err error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
	)
	if query, err = pixiv.NewQuery(map[string]interface{}{"lang": "jp", "full": 1}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/user", pid), query, nil)
	if err != nil {
		return nil, err
	}

	body = &dtos.UserInfoDTO{}
	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func GetProfileAll(ctx pixiv.IContext, pid int64) (_ *dtos.AllProfileDTO, err error) {
	var (
		c    pixiv.IClient
		body dtos.AllProfileDTO
	)
	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/user/", pid, "/profile", All), nil, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func GetProfileTop(ctx pixiv.IContext, pid int64) (_ *dtos.TopProfileDTO, err error) {
	var (
		c    pixiv.IClient
		body dtos.TopProfileDTO
	)
	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/user/", pid, "/profile", Top), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	return &body, nil
}

func GetFanboxUlr(ctx pixiv.IContext, pid int64) (u string, err error) {
	var c pixiv.IClient
	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return "", err
	}

	header, err := pixiv.Header(c.Get, pixiv.Path("/fanbox/creator/", pid), nil, nil)
	if err != nil {
		return "", err
	}
	return header.Get("location"), nil
}
