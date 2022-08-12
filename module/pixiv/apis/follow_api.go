package apis

import (
	"encoding/json"
	"net/http"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

type FollowAPI struct{}

func followLast(ctx pixiv.Context, mode string, page int) (*FollowLatestDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"p":    page,
		"mode": mode,
		"lang": "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Json(ctx, http.MethodGet, pixiv.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return nil, err
	}

	body := &FollowLatestDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (FollowAPI) Illusts(ctx pixiv.Context, page int) (*FollowLatestDTO, error) {
	return followLast(ctx, Illust, page)
}

func (FollowAPI) Novel(ctx pixiv.Context, page int) (*FollowLatestDTO, error) {
	return followLast(ctx, Novel, page)
}

func following(ctx pixiv.Context, uid int32, rest string, limit int, offset int) (*FollowDTO, error) {
	// 添加参数
	query, err := pixiv.NewQuery(map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"rest":   rest,
		"tag":    "",
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Json(ctx, http.MethodGet, pixiv.Path("/ajax/user", uid, "/following"), query, nil)
	if err != nil {
		return nil, err
	}

	body := FollowDTO{}
	if err := json.Unmarshal(data, &body); err != nil {
		return nil, err
	}
	return &body, nil
}

// Following return ShowFollowing User
func (FollowAPI) Following(ctx pixiv.Context, uid int32, limit int, offset int) (*FollowDTO, error) {
	return following(ctx, uid, Show, limit, offset)
}
