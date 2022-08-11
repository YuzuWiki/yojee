package pixivsrv

import (
	"fmt"
	"github.com/neilotoole/errgroup"
	"golang.org/x/net/context"
	"strconv"

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

func NewService(phpSessID string, numG, qSize int) Service {
	taskG, ch := errgroup.WithContextN(context.Background(), numG, qSize)

	return Service{
		ctx:       pixiv.NewContext(phpSessID),
		taskGroup: taskG,
		taskCtx:   ch,
	}
}

func (srv *Service) SyncUser(pid int64) error {
	info, err := srv.apiInfo.Information(srv.ctx, pid)
	if err != nil {
		return err
	}

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

	for _artId := range artwork.Illusts {
		if artId, err := strconv.ParseInt(_artId, 10, 0); err != nil {
			srv.asyncArt(apis.Illust, artId)
		}
	}

	for _artId := range artwork.Manga {
		if artId, err := strconv.ParseInt(_artId, 10, 0); err != nil {
			srv.asyncArt(apis.Manga, artId)
		}
	}

	for _artId := range artwork.Novel {
		if artId, err := strconv.ParseInt(_artId, 10, 0); err != nil {
			srv.asyncArt(apis.Novel, artId)
		}
	}
	return nil
}

func (srv *Service) Wait() error {
	return srv.taskGroup.Wait()
}
