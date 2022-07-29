package user

import (
	"encoding/json"
	"net/http"

	pixivService "github.com/like9th/yojee/yojee/service/PixivService"
)

type InfoAPI struct{}

func (api InfoAPI) Extra(ctx pixivService.ContextVar) (*ExtraDTO, error) {
	data, err := pixivService.Request(ctx, http.MethodGet, pixivService.Path("/ajax/user", "extra"), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ExtraDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
