package apis

import (
	"encoding/json"
	pixiv2 "github.com/YuzuWiki/yojee/module/pixiv"
	"net/http"
)

type FollowAPI struct{}

func followLast(ctx pixiv2.Context, mode string, page int) (*FollowLatestDTO, error) {
	query, err := pixiv2.NewQuery(map[string]interface{}{
		"p":    page,
		"mode": mode,
		"lang": "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv2.Request(ctx, http.MethodGet, pixiv2.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return nil, err
	}

	body := &FollowLatestDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (api FollowAPI) Illusts(ctx pixiv2.Context, page int) (*FollowLatestDTO, error) {
	return followLast(ctx, Illust, page)
}

func (api FollowAPI) Novel(ctx pixiv2.Context, page int) (*FollowLatestDTO, error) {
	return followLast(ctx, Novel, page)
}

func following(ctx pixiv2.Context, uid int32, rest string, limit int, offset int) (*FollowDTO, error) {
	// 添加参数
	query, err := pixiv2.NewQuery(map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"rest":   rest,
		"tag":    "",
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv2.Request(ctx, http.MethodGet, pixiv2.Path("/ajax/user", uid, "/following"), query, nil)
	if err != nil {
		return nil, err
	}

	body := FollowDTO{}
	if err := json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (api FollowAPI) FollowingUsers(ctx pixiv2.Context, uid int32, limit int, offset int, isShow bool) (*FollowDTO, error) {
	if isShow {
		return following(ctx, uid, Show, limit, offset)
	}
	return following(ctx, uid, Hide, limit, offset)
}
