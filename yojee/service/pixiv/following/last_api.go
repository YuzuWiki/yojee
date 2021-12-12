package following

import (
	"context"
	"encoding/json"

	"github.com/like9th/yojee/yojee/service/pixiv"
)

type LastAPI struct{}

func (api LastAPI) FollowLatestIllust(ctx context.Context, mold string, mode string, page int) (*FollowLatestDTO, error) {
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
