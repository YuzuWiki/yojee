package user

import (
	"github.com/like9th/yojee/yojee/common/requests"
	"io/ioutil"

	"github.com/like9th/yojee/yojee/service/PixivService"
)

type BookMarkAPI struct{}

func (api BookMarkAPI) get(ctx pixivService.ContextVar, u string, query *requests.Query, params *requests.Params) ([]byte, error) {
	resp, err := ctx.Client().Get(u, query, params)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer resp.Body.Close()

	return ioutil.ReadAll(body) // TODO 可能调整为 struct
}
