package global

import (
	"github.com/YuzuWiki/yojee/module/pixiv_v2"
	"github.com/YuzuWiki/yojee/module/pixiv_v2/requests"
)

var Pixiv pixiv_v2.ISession

func init() {
	if Pixiv == nil {
		Pixiv = requests.NewSession()
	}
}
