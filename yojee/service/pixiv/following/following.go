package following

import (
	"context"
	"encoding/json"
	"github.com/like9th/yojee/yojee/service/pixiv"
)

type FollowAPI struct{}

func (api FollowAPI) get(ctx context.Context, uid int32, rest string, limit int, offset int) (*FollowDTO, error) {
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

	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user", uid, "following"), query)
	if err != nil {
		return nil, err
	}

	data := FollowDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (api FollowAPI) FindShow(ctx context.Context, uid int32, limit int, offset int) (*FollowDTO, error) {
	return api.get(ctx, uid, StatusShow, limit, offset)
}

func (api FollowAPI) FindHide(ctx context.Context, uid int32, limit int, offset int) (*FollowDTO, error) {
	return api.get(ctx, uid, StatusHide, limit, offset)
}
