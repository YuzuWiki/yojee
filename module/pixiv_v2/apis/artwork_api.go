package apis

import (
	"encoding/json"
	"strconv"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/dtos"
)

func GetAccountPid(ctx pixiv_v2.IContext) (int64, error) {
	c, err := global.Pixiv.New(ctx.PhpSessID())
	if err != nil {
		return 0, err
	}

	resp, err := c.Get("https://"+pixiv_v2.PixivHost, nil, nil)
	if err != nil {
		return 0, err
	}

	if pid := resp.Header.Get("x-userid"); len(pid) > 0 {
		return strconv.ParseInt(pid, 10, 64)
	}
	return 0, nil
}

func getArtWork(ctx pixiv_v2.IContext, artType string, artId int64) (body *dtos.ArtworkDTO, err error) {
	var (
		query *pixiv_v2.Query
		c     pixiv_v2.IClient
	)

	if query, err = pixiv_v2.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return
	}

	if c, err = global.Pixiv.New(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := pixiv_v2.Json(c.Get, pixiv_v2.Path("/ajax", artType, artId), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}

	return body, nil
}

func GetIllusts(ctx pixiv_v2.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func GetMangas(ctx pixiv_v2.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func GetNovels(ctx pixiv_v2.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Novel, artId)
}
