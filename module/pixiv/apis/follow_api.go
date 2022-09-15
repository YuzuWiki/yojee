package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func followArtWork(ctx pixiv.IContext, mode string, page int) (body *dtos.FollowLatestDTO, err error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
	)

	if query, err = pixiv.NewQuery(map[string]interface{}{"p": page, "mode": mode, "lang": "jp"}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func FollowIllusts(ctx pixiv.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Illust, page)
}

func FollowNovel(ctx pixiv.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Novel, page)
}

func FollowManga(ctx pixiv.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Manga, page)
}
