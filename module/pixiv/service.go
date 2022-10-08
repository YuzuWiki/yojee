package pixiv

import (
	"os"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
)

var DefaultContext common.IContext

type Service struct{}

func init() {
	// inti default context
	DefaultContext = Pixivlee.NewContext(os.Getenv("PIXIV_PHPSESSID"))

	// init proxy
	Pixivlee.SetGlobalProxy(os.Getenv("HTTP_PROXY"))
}
