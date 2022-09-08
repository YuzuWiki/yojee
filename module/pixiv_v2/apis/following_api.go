package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/dtos"
)

func FollowingArtWork(ctx pixiv_v2.IContext, uid int32, limit int32, offset int) (body *dtos.FollowingDTO, err error) {
	var (
		query *pixiv_v2.Query
		c     pixiv_v2.IClient
	)

	if query, err = pixiv_v2.NewQuery(map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"rest":   Show,
		"tag":    "",
		"lang":   "zh",
	}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax/user", uid, "/following"), query, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}

	return body, nil
}
