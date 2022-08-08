package pixivsrv

import (
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
	"github.com/YuzuWiki/yojee/service/modelsrv"
)

type Service struct {
	ctx pixiv.Context

	// srv apis
	infoAPi apis.InfoAPI

	// srv mod
	pixivSrv modelsrv.PixivUser
}

func NewService(phpSessID string) Service {
	return Service{
		ctx: pixiv.NewContext(phpSessID),
	}
}

func (srv *Service) syncUserinfo(pid int64) error {
	info, err := srv.infoAPi.Info(srv.ctx, pid)
	if err != nil {
		return err
	}

	if err := srv.pixivSrv.InsertUser(*info); err != nil {
		return err
	}
	return nil
}
