package apis

import (
	"encoding/json"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"net/http"
)

type ArtworkAPI struct{}

func (ArtworkAPI) Illust(ctx pixiv.Context, artId int64) (*ArtworkIllustDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"lang": "jp",
	})
	if err != nil {
		return nil, err
	}
	// https://www.pixiv.net/ajax/illust/90735220?lang=jp
	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax/illust", artId), query, nil)
	if err != nil {
		return nil, err
	}

	body := &ArtworkIllustDTO{}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}
