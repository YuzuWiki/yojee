package apis

import (
	"encoding/json"
	"net/http"

	"yojee/service/pixiv"
)

type ProfileAPI struct{}

func (api ProfileAPI) All(ctx pixiv.ContextVar, uid int64) (*ProfileAllDTO, error) {
	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user/", uid, "/profile", All), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ProfileAllDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, err
}

func (api ProfileAPI) Top(ctx pixiv.ContextVar, uid int64) (*ProfileTopDTO, error) {
	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user/", uid, "/profile", Top), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ProfileTopDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, err
}
