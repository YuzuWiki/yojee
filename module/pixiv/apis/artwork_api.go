package apis

import (
	"encoding/json"
	"strconv"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func GetAccountPid(ctx pixiv.IContext) (int64, error) {
	c, err := global.Pixiv.New(ctx.PhpSessID())
	if err != nil {
		return 0, err
	}

	header, err := pixiv.Header(c.Get, "https://"+pixiv.PixivHost, nil, nil)
	if err != nil {
		return 0, err
	}

	if pid := header.Get("x-userid"); len(pid) > 0 {
		return strconv.ParseInt(pid, 10, 64)
	}
	return 0, nil
}

func getArtWork(ctx pixiv.IContext, artType string, artId int64) (body *dtos.ArtworkDTO, err error) {
	var (
		query *pixiv.Query
		c     pixiv.IClient
	)

	if query, err = pixiv.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv.Json(c.Get, pixiv.Path("/ajax", artType, artId), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}

	return body, nil
}

func GetIllusts(ctx pixiv.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func GetMangas(ctx pixiv.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Manga, artId)
}

func GetNovels(ctx pixiv.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Novel, artId)
}
