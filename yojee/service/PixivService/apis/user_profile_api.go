package apis

import (
	"encoding/json"
	"net/http"

	pixivService "github.com/like9th/yojee/yojee/service/PixivService"
)

type ProfileAPI struct{}

func (api ProfileAPI) All(ctx pixivService.ContextVar, uid int64) (*ProfileAllDTO, error) {
	data, err := pixivService.Request(ctx, http.MethodGet, pixivService.Path("/ajax/user/", uid, "/profile", pixivService.All), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ProfileAllDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, err
}

func (api ProfileAPI) Top(ctx pixivService.ContextVar, uid int64) (*ProfileTopDTO, error) {
	data, err := pixivService.Request(ctx, http.MethodGet, pixivService.Path("/ajax/user/", uid, "/profile", pixivService.Top), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ProfileTopDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, err
}
