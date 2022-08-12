package pixivsrv

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/neilotoole/errgroup"
	"golang.org/x/net/context"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
	"github.com/YuzuWiki/yojee/service/modelsrv"
)

type Service struct {
	ctx pixiv.Context

	// task channel
	taskCtx context.Context

	// task taskGroup
	taskGroup *errgroup.Group

	// srv apis
	apiInfo    apis.InfoAPI
	apiArtwork apis.ArtworkAPI

	// srv mod
	modUser    modelsrv.PixivUser
	modTag     modelsrv.PixivTag
	modArtwork modelsrv.PixivArtwork
}

func (srv *Service) SyncUser(pid int64) error {
	info, err := srv.apiInfo.Information(srv.ctx, pid)
	if err != nil {
		return err
	}
	global.Logger.Info().Msg(fmt.Sprintf("[SyncUser] info: pid=%d, data=%+v", pid, info))

	if err := srv.modUser.InsertUser(*info); err != nil {
		return err
	}
	return nil
}

func (srv *Service) asyncArtTag(artType string, artId int64, tagName, romaji string) func() error {
	return func() error {
		tagId, err := srv.modTag.Insert(tagName, romaji)
		if err != nil {
			return err
		}

		return srv.modArtwork.MarkTag(artType, artId, tagId)
	}
}

func (srv *Service) asyncArt(artType string, artId int64) func() error {
	return func() error {
		var fn func(pixiv.Context, int64) (*apis.ArtworkDTO, error)
		switch artType {
		case apis.Illust:
			fn = srv.apiArtwork.Illust
		case apis.Manga:
			fn = srv.apiArtwork.Manga
		case apis.Novel:
			fn = srv.apiArtwork.Novel
		default:
			return fmt.Errorf(fmt.Sprintf("Unsuport ArtType (%s)", artType))
		}

		artwork, err := fn(srv.ctx, artId)
		if err != nil {
			return err
		}
		global.Logger.Info().Msg(fmt.Sprintf("[asyncArt] artwork: artId=%d data=%+v", artId, artwork))

		if _, err := srv.modArtwork.Insert(*artwork); err != nil {
			return err
		}

		for _, tag := range artwork.Tags.Tags {
			srv.taskGroup.Go(srv.asyncArtTag(artType, artwork.ArtId, tag.Name, tag.Romaji))
		}
		return nil
	}
}

func (srv *Service) SyncArtworks(pid int64) error {
	artwork, err := srv.apiInfo.Artwork(srv.ctx, pid)
	if err != nil {
		return err
	}
	global.Logger.Info().Msg(fmt.Sprintf("[SyncArtworks] artwork: pid=%d data=%+v", pid, artwork))

	for _artId := range artwork.Illusts {
		global.Logger.Info().Msg(fmt.Sprintf("[SyncArtworks] artwork: _artId=%s", _artId))
		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
			global.Logger.Info().Msg(fmt.Sprintf("[SyncArtworks] artwork: artId=%d", artId))

			srv.taskGroup.Go(srv.asyncArt(apis.Illust, artId))
		}
	}

	for _artId := range artwork.Manga {
		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
			srv.taskGroup.Go(srv.asyncArt(apis.Manga, artId))
		}
	}

	for _artId := range artwork.Novel {
		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
			srv.taskGroup.Go(srv.asyncArt(apis.Novel, artId))
		}
	}
	return nil
}

type ServiceInterface interface {
	SyncUser(int64) error
	SyncArtworks(int64) error
}

var pMu sync.Mutex
var _services map[string]ServiceInterface

func NewService(phpSessID string, numG int, qSize int) ServiceInterface {
	if srv, isOk := _services[phpSessID]; isOk {
		return srv
	}

	pMu.Lock()
	defer pMu.Unlock()

	taskG, ch := errgroup.WithContextN(context.Background(), numG, qSize)
	srv := Service{
		ctx:       pixiv.NewContext(phpSessID),
		taskGroup: taskG,
		taskCtx:   ch,
	}
	_services[phpSessID] = &srv

	return &srv
}

func init() {
	_services = map[string]ServiceInterface{}
}
