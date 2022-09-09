package pixivsrv

// import (
// 	"fmt"
// 	"strconv"
//
// 	"github.com/panjf2000/ants"
//
// 	"github.com/YuzuWiki/yojee/global"
// 	"github.com/YuzuWiki/yojee/module/pixiv"
// 	"github.com/YuzuWiki/yojee/module/pixiv/apis"
// 	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
// 	"github.com/YuzuWiki/yojee/service/modelsrv"
// )
//
// type Service struct {
// 	ctx pixiv.Context
//
// 	//  task taskGroup
// 	jobPool *ants.Pool
//
// 	//  srv apis
// 	apiInfo    apis.InfoAPI
// 	apiArtwork apis.ArtworkAPI
//
// 	//  srv mod
// 	modUser    modelsrv.PixivUser
// 	modTag     modelsrv.PixivTag
// 	modArtwork modelsrv.PixivArtwork
// }
//
// func (srv *Service) SyncUser(pid int64) error {
// 	info, err := srv.apiInfo.Information(srv.ctx, pid)
// 	if err != nil {
// 		return err
// 	}
// 	global.Logger.Info().Msg(fmt.Sprintf("[SyncUser] info: pid=%d, data=%+v", pid, info))
//
// 	if err := srv.modUser.InsertUser(*info); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (srv *Service) asyncArtTag(artType string, artId int64, jp, romaji string) func() {
// 	return func() {
// 		tagId, err := srv.modTag.Insert(jp, romaji)
// 		if err != nil {
// 			global.Logger.Error().Msg(fmt.Sprintf("[AsyncArtTag] error: art_type=%s art_id=%d tag_name=%s errmsg=%s", artType, artId, jp, err.Error()))
// 			return
// 		}
//
// 		if err := srv.modArtwork.MarkTag(artType, artId, tagId); err != nil {
// 			global.Logger.Error().Msg(fmt.Sprintf("[AsyncArtTag] error: art_type=%s art_id=%d tag_name=%s errmsg=%s", artType, artId, jp, err.Error()))
// 		}
// 	}
// }
//
// func (srv *Service) asyncArt(artType string, artId int64) func() {
// 	return func() {
// 		var fn func(pixiv.Context, int64) (*apis.ArtworkDTO, error)
// 		switch artType {
// 		case apis.Illust:
// 			fn = srv.apiArtwork.Illust
// 		case apis.Manga:
// 			fn = srv.apiArtwork.Manga
// 		case apis.Novel:
// 			fn = srv.apiArtwork.Novel
// 		default:
// 			global.Logger.Error().Msg(fmt.Sprintf("[AsyncArt] error: art_type=%s art_id=%d errmsg=Unsuport ArtType", artType, artId))
// 			return
// 		}
//
// 		artwork, err := fn(srv.ctx, artId)
// 		if err != nil {
// 			global.Logger.Error().Msg(fmt.Sprintf("[AsyncArt] apiArtwork error: art_type=%s art_id=%d errmsg=%s", artType, artId, err.Error()))
// 			return
// 		}
//
// 		if _, err := srv.modArtwork.Insert(*artwork); err != nil {
// 			global.Logger.Error().Msg(fmt.Sprintf("[AsyncArt] insert error: art_type=%s art_id=%d errmsg=%s", artType, artId, err.Error()))
// 			return
// 		}
//
// 		for _, tag := range artwork.Tags.Tags {
// 			if err := srv.jobPool.Submit(srv.asyncArtTag(artType, artwork.ArtId, tag.Jp, tag.Romaji)); err != nil {
// 				global.Logger.Error().Msg(fmt.Sprintf("[AsyncArt] Submit AsyncArtTag error: art_type=%s art_id=%d errmsg=%s", artType, artId, err.Error()))
// 			}
// 		}
// 		return
// 	}
// }
//
// func (srv *Service) SyncArtworks(pid int64) error {
// 	artwork, err := srv.apiInfo.Artwork(srv.ctx, pid)
// 	if err != nil {
// 		return err
// 	}
// 	global.Logger.Info().Msg(fmt.Sprintf("[SyncArtworks] artwork: pid=%d data=%+v", pid, artwork))
//
// 	for _artId := range artwork.Illusts {
// 		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
// 			if err := srv.jobPool.Submit(srv.asyncArt(apis.Illust, artId)); err != nil {
// 				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtworks] Submit AsyncArt error: art_type=%s art_id=%d errmsg=%s", apis.Illust, artId, err.Error()))
// 			}
// 		}
// 	}
//
// 	for _artId := range artwork.Manga {
// 		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
// 			if err := srv.jobPool.Submit(srv.asyncArt(apis.Manga, artId)); err != nil {
// 				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtworks] Submit AsyncArt error: art_type=%s art_id=%d errmsg=%s", apis.Illust, artId, err.Error()))
// 			}
// 		}
// 	}
//
// 	for _artId := range artwork.Novel {
// 		if artId, err := strconv.ParseInt(_artId, 10, 0); err == nil {
// 			if err := srv.jobPool.Submit(srv.asyncArt(apis.Novel, artId)); err != nil {
// 				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtworks] Submit AsyncArt error: art_type=%s art_id=%d errmsg=%s", apis.Illust, artId, err.Error()))
// 			}
// 		}
// 	}
// 	return nil
// }
//
// func NewService(phpSessID string, numG int, qSize int) Service {
// 	pool, _ := ants.NewPool(200)
// 	pool.MaxBlockingTasks = 5
// 	return Service{
// 		ctx:     global.NewContext(phpSessID),
// 		jobPool: pool,
// 	}
// }
