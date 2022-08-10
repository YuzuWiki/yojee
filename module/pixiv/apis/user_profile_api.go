package apis

import (
	"encoding/json"
	"net/http"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

type ProfileAPI struct{}

func (api ProfileAPI) All(ctx pixiv.Context, uid int64) (*ProfileAllDTO, error) {
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

func (api ProfileAPI) Top(ctx pixiv.Context, uid int64) (*ProfileTopDTO, error) {
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
