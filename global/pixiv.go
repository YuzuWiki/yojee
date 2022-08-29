package global

import "github.com/YuzuWiki/yojee/module/pixiv_v2"

var Pixiv pixiv_v2.ISession

func init() {
	if Pixiv == nil {
		Pixiv = pixiv_v2.NewSession()
	}
}
