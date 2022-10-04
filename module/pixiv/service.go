package pixiv

import (
	"os"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
)

var DefaultContext common.IContext

type Service struct{}

func init() {
	sessId := os.Getenv("PIXIV_PHPSESSID")
	proxyUrl := os.Getenv("PROXY_URL")

	// inti default context
	DefaultContext = Pixivlee.NewContext(sessId)

	// init proxy
	if c, err := Pixivlee.Pool().Get(sessId); err == nil && len(proxyUrl) > 0 {
		if err = c.SetProxy(proxyUrl); err != nil {
			panic(err.Error())
		}
	}
}
