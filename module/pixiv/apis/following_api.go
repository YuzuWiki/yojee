package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func GetFollowing(ctx pixiv.IContext, pid int64, limit int, offset int) (*dtos.FollowingDTO, error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
		err   error
	)

	if query, err = pixiv.NewQuery(map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"rest":   Show,
		"tag":    "",
		"lang":   "jp",
	}); err != nil {
		return nil, err
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/user", pid, "/following"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &dtos.FollowingDTO{}
	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
