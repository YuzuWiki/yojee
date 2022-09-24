package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func (Service) GetArtwork(artType string, artId int64) (*ArtworkDTO, error) {
	var artworks []ArtworkDTO
	if err := global.DB().Raw(`
	 SELECT
		artwork.*,
		(
			SELECT
				GROUP_CONCAT(tag.jp)
			FROM pixiv_tag          AS tag
			JOIN pixiv_artwork_tag  AS pat
				ON pat.tag_id    = tag.tag_id
			WHERE pat.is_deleted = false
			  AND pat.art_type   = artwork.art_type
			  AND pat.art_id     = artwork.art_id
		) AS tags
	FROM pixiv_artwork AS artwork
	WHERE artwork.art_type = ?
	  AND artwork.art_id = ?
	LIMIT 1;`, artType, artId,
	).Find(&artworks).Error; err != nil {
		return nil, err
	}

	if len(artworks) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return &artworks[0], nil
}

func insertArtwork(artType string, data *dtos.ArtworkDTO) (int64, error) {
	row := model.PixivArtworkMod{
		Pid:           data.Pid,
		ArtId:         data.ArtId,
		ArtType:       artType,
		Title:         data.Title,
		Description:   data.Description,
		ViewCount:     data.ViewCount,
		LikeCount:     data.LikeCount,
		BookmarkCount: data.BookmarkCount,
		CreateDate:    &data.CreateDate,
	}
	if err := global.DB().FirstOrCreate(&row, model.PixivArtworkMod{Pid: data.Pid, ArtType: artType, ArtId: data.ArtId}).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}

func SyncArtWork(artType string, artId int64) (err error) {
	var (
		artwork *dtos.ArtworkDTO
		tagId   int64
	)

	// 获取作品信息
	global.JobPool.Submit(func() { artwork, err = apis.GetIllusts(pixiv.DefaultContext, artId) })
	if err != nil {
		return err
	}

	if _, err = insertArtwork(artType, artwork); err != nil {
		return err
	}

	for _, tag := range artwork.Tags.Tags {
		global.JobPool.Submit(func() { tagId, err = SyncTag(tag.Jp) })
		if err != nil {
			continue
		}

		if err = markArtworkTag(artType, artId, tagId); err != nil {
			continue
		}
	}
	return nil
}

func (s Service) SyncArtWorks(pid int64) (err error) {
	if _, err = s.FlushAccountInfo(pid); err != nil {
		return err
	}

	var profile *dtos.AllProfileDTO
	global.JobPool.Submit(func() { profile, err = apis.GetProfileAll(pixiv.DefaultContext, pid) })
	if err != nil {
		return err
	}

	for artType, artIds := range map[string]dtos.ArtWorkIdsDTO{
		apis.Illust: profile.Illusts,
		apis.Manga:  profile.Manga,
		apis.Novel:  profile.Novel,
	} {
		for idx, artId := range artIds {
			if err = SyncArtWork(artType, artId); err != nil {
				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork] SyncArtWork ERROR: (%d) artType=%s artId=%d  errmsg=%s", pid, artType, artId, err.Error()))
				return err
			}
			global.Logger.Debug().Msg(fmt.Sprintf("[SyncArtWork] SyncArtWork RUNNING: (%d) artType=%s artId=%d  %d/%d ", pid, artType, artId, idx+1, len(artIds)))
		}
	}
	return nil
}
