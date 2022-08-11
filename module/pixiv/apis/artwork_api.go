package apis

import (
	"encoding/json"
	"net/http"

	"github.com/YuzuWiki/yojee/module/pixiv"
)

type ArtworkAPI struct{}

func getArtWork(ctx pixiv.Context, artType string, artId int64) (*ArtworkDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"lang": "jp",
	})
	if err != nil {
		return nil, err
	}

	data, err := pixiv.Request(ctx, http.MethodGet, pixiv.Path("/ajax", artType, artId), query, nil)
	if err != nil {
		return nil, err
	}

	body := &ArtworkDTO{ArtType: artType}
	if err := json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (ArtworkAPI) Illust(ctx pixiv.Context, artId int64) (*ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func (ArtworkAPI) Manga(ctx pixiv.Context, artId int64) (*ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func (ArtworkAPI) Novel(ctx pixiv.Context, artId int64) (*ArtworkDTO, error) {
	return getArtWork(ctx, Novel, artId)
}
