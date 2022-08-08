package pixiv

import (
	"github.com/YuzuWiki/yojee/module/pixiv"
)

type Service struct {
	ctx pixiv.Context
}

func NewService(phpSessID string) Service {
	return Service{
		ctx: pixiv.NewContext(phpSessID),
	}
}
