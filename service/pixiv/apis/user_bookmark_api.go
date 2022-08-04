package apis

import (
	"encoding/json"
	"net/http"

	"yojee/service/pixiv"
)

type BookMarkAPI struct{}

func (api BookMarkAPI) FindShow(ctx pixiv.ContextVar, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   Show,
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &BookmarkDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (api BookMarkAPI) FindHide(ctx pixiv.ContextVar, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   Hide,
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &BookmarkDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (api BookMarkAPI) GetIllustTags(ctx pixiv.ContextVar, uid int64, tag string, offset int, limit int) (*BookMarkTagsDTO, error) {
	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user", uid, "/illusts/bookmark/tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &BookMarkTagsDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}
