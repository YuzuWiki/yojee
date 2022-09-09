package global

import (
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/requests"
)

var Pixiv pixiv.ISession

func initPixiv() {
	if Pixiv == nil {
		Pixiv = requests.NewSession()
	}
}
