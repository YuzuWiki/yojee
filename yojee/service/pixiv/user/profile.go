package user

import (
	"context"
	"encoding/json"
	"github.com/like9th/yojee/yojee/service/pixiv"
)

/*

收藏: https://www.pixiv.net/ajax/user/20376220/illusts/bookmarks?tag=&offset=0&limit=48&rest=show&lang=zh
tag标签: https://www.pixiv.net/ajax/user/20376220/illusts/bookmark/tags?lang=zh
*/

type ProfileAPI struct {}

func (api ProfileAPI) All(ctx context.Context,  userID int) (*ProfileAllDTO, error) {
	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user/", userID, "/profile", ModAll), nil)
	if err != nil {
		return nil, err
	}

	data := ProfileAllDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, err
}


func (api ProfileAPI) Top(ctx context.Context,  userID int) (*ProfileTopDTO, error) {
	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user/", userID, "/profile", ModTop), nil)
	if err != nil {
		return nil, err
	}

	data := ProfileTopDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, err
}

