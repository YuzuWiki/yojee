package apis

import (
	"golang.org/x/net/context"
	"strconv"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
)

func GetAccountPid(PhpSessID string) (int64, error) {
	c, err := global.Pixiv.New(PhpSessID)
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

func Get(ctx context.Context) {

}
