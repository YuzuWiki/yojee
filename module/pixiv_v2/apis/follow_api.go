package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/dtos"
)

func followArtWork(ctx pixiv_v2.IContext, mode string, page int) (body *dtos.FollowLatestDTO, err error) {
	var (
		query *pixiv_v2.Query
		c     pixiv_v2.IClient
	)

	if query, err = pixiv_v2.NewQuery(map[string]interface{}{"p": page, "mode": mode, "lang": "zh"}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func FollowIllusts(ctx pixiv_v2.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Illust, page)
}

func FollowNovel(ctx pixiv_v2.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Novel, page)
}

func FollowManga(ctx pixiv_v2.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Manga, page)
}
