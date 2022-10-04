package pixiv

import (
	"os"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
)

var DefaultContext common.IContext

type Service struct{}

func init() {
	sessid := os.Getenv("PIXIV_PHPSESSID")

	// inti default context
	DefaultContext = Pixivlee.NewContext(sessid)

	// init proxy
	if c, err := Pixivlee.Pool().Get(os.Getenv("PIXIV_PHPSESSID")); err == nil {
		if err = c.SetProxy(os.Getenv("PROXY_URL")); err != nil {
			panic(err.Error())
		}
	}
}
