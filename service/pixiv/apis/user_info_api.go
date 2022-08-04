package apis

import (
	"encoding/json"
	"net/http"

	"github.com/YuzuWiki/yojee/service/pixiv"
)

type InfoAPI struct{}

func (api InfoAPI) Extra(ctx pixiv.ContextVar) (*ExtraDTO, error) {
	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/user", "extra"), nil, nil)
	if err != nil {
		return nil, err
	}

	body := &ExtraDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
