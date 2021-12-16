package user

import (
	"context"
	"encoding/json"
	"github.com/like9th/yojee/yojee/service/pixiv"
)

type SelfAPI struct{}

func (api *SelfAPI) Extra(ctx context.Context) (*ExtraDTO, error) {
	body, err := pixiv.Get(ctx, pixiv.Path("/ajax/user", "extra"), nil)
	if err != nil {
		return nil, err
	}

	data := ExtraDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, err
}
