package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func bookmark(ctx pixiv.IContext, rest string, uid int64, tag string, offset int, limit int) (body *dtos.BookmarkDTO, err error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
	)
	if query, err = pixiv.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   rest,
		"lang":   "jp",
	}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}
	return body, err
}

func BookmarkShow(ctx pixiv.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Show, uid, tag, offset, limit)
}

func BookmarkHide(ctx pixiv.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Hide, uid, tag, offset, limit)
}
