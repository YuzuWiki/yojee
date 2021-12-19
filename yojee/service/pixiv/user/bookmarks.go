package user

import (
	"context"
	"encoding/json"

	"github.com/like9th/yojee/yojee/service/pixiv"
)

type BookMarkAPI struct{}

func (api BookMarkAPI) get(ctx context.Context, uid int64, tag string, offset int, limit int, rest string) (*BookmarkDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   rest,
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user", uid, "/illusts/bookmarks"), query)
	if err != nil {
		return nil, err
	}

	data := BookmarkDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (api BookMarkAPI) FindShow(ctx context.Context, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	return api.get(ctx, uid, tag, offset, limit, StatusShow)
}

func (api BookMarkAPI) FindHide(ctx context.Context, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	return api.get(ctx, uid, tag, offset, limit, StatusHide)
}

func (api BookMarkAPI) GetIllustTags(ctx context.Context, uid int64) (*BookMarkTagsDTO, error) {
	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user", uid, "/illusts/bookmark/tags"), nil)
	if err != nil {
		return nil, err
	}

	data := BookMarkTagsDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
