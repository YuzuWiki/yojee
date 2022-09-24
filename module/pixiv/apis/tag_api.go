package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func GetTag(ctx pixiv.IContext, jP string) (*dtos.TagDTO, error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
		err   error
		tag   dtos.TagDTO
	)

	if query, err = pixiv.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return nil, err
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax/search/tags", jP), query, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}
