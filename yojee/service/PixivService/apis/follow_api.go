package apis

import (
	"encoding/json"
	"net/http"

	pixivService "github.com/like9th/yojee/yojee/service/PixivService"
)

type FollowAPI struct{}

/*
	LastIllusts return data & error
		@param

*/
func (api FollowAPI) LastIllusts(ctx pixivService.ContextVar, mode string, page int) (*FollowLatestDTO, error) {
	query, err := pixivService.NewQuery(map[string]interface{}{
		"p":    page,
		"mode": mode,
		"lang": "zh",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixivService.Request(ctx, http.MethodGet, pixivService.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return nil, err
	}

	body := &FollowLatestDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
