package following

import (
	"context"
	"encoding/json"

	"github.com/like9th/yojee/yojee/service/pixiv"
)

type FollowLastAPI struct{}

func (api FollowLastAPI) get(ctx context.Context, mold string, mode string, page int) (*FollowLatestDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"p":    page,
		"mode": mode,
		"lang": "zh",
	})
	if err != nil {
		return nil, err
	}

	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/follow_latest", mold), query)
	if err != nil {
		return nil, err
	}

	data := FollowLatestDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (api FollowLastAPI) Illust(ctx context.Context, mode string, page int) (*FollowLatestDTO, error) {
	return api.get(ctx, Mold_Illust, mode, page)
}

func (api FollowLastAPI) Novel(ctx context.Context, mode string, page int) (*FollowLatestDTO, error) {
	return api.get(ctx, Mold_Novel, mode, page)
}
