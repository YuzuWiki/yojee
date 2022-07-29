package user

import (
	"encoding/json"
	"github.com/like9th/yojee/yojee/common/requests"
	"io/ioutil"

	"github.com/like9th/yojee/yojee/service/PixivService"
)

type BookMarkAPI struct{}

func (api BookMarkAPI) get(ctx pixivService.ContextVar, u string, query *requests.Query, params *requests.Params) ([]byte, error) {
	resp, err := ctx.Client().Get(u, query, params)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer resp.Body.Close()

	return ioutil.ReadAll(body)
}

func (api BookMarkAPI) FindShow(ctx pixivService.ContextVar, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	query, err := pixivService.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   pixivService.Show,
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := api.get(ctx, pixivService.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &BookmarkDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (api BookMarkAPI) FindHide(ctx pixivService.ContextVar, uid int64, tag string, offset int, limit int) (*BookmarkDTO, error) {
	query, err := pixivService.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   pixivService.Hide,
		"lang":   "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := api.get(ctx, pixivService.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &BookmarkDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}

func (api BookMarkAPI) GetIllustTags(ctx pixivService.ContextVar, uid int64, tag string, offset int, limit int) (*BookMarkTagsDTO, error) {
	data, err := api.get(ctx, pixivService.Path("/ajax/user", uid, "/illusts/bookmark/tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &BookMarkTagsDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}

	return body, nil
}