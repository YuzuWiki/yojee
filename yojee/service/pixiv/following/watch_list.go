package following

import (
	"context"
	"encoding/json"
	"github.com/like9th/yojee/yojee/service/pixiv"
)

/*
WatchListAPI
追更列表中的作品:
	漫画: https://www.pixiv.net/ajax/watch_list/manga?p=1&lang=zh
	小说: https://www.pixiv.net/ajax/watch_list/novel?p=1&new=1&lang=zh
*/
type WatchListAPI struct{}

func (api WatchListAPI) get(ctx context.Context, mold string, page int) (*WatchListDTO, error) {
	query, err := pixiv.NewQuery(map[string]interface{}{
		"p":    page,
		"lang": "zh",
	})
	if err != nil {
		return nil, err
	}

	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/watch_list", mold), query)
	if err != nil {
		return nil, err
	}

	data := WatchListDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (api WatchListAPI) Manga(ctx context.Context, page int) (*WatchListDTO, error) {
	return api.get(ctx, Mold_Manga, page)
}

func (api WatchListAPI) Novel(ctx context.Context, page int) (*WatchListDTO, error) {
	return api.get(ctx, Mold_Novel, page)
}
