package apis

import (
	"encoding/json"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"net/http"
)

type ArtworkAPI struct{}

func (ArtworkAPI) Illust(ctx pixiv.Context, artId int64) (*ArtworkDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"lang": "jp",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/illust", artId), query, nil)
	if err != nil {
		return nil, err
	}

	body := &ArtworkDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
