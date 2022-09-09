package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/dtos"
)

func bookmark(ctx pixiv_v2.IContext, rest string, uid int64, tag string, offset int, limit int) (body *dtos.BookmarkDTO, err error) {
	var (
		query *pixiv_v2.Query
		c     pixiv_v2.IClient
	)
	if query, err = pixiv_v2.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   rest,
		"lang":   "zh",
	}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}
	return body, err
}

func BookmarkShow(ctx pixiv_v2.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Show, uid, tag, offset, limit)
}

func BookmarkHide(ctx pixiv_v2.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Hide, uid, tag, offset, limit)
}
