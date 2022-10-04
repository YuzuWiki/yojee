package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/dtos"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
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

func (Service) GetArtworks(pid int64, artType string, limit, offset int) (*[]ArtworkDTO, error) {
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
	  AND artwork.pid = ?
	LIMIT ? OFFSET ?;`, artType, pid, limit, offset,
	).Find(&artworks).Error; err != nil {
		return nil, err
	}

	return &artworks, nil
}

func SyncArtWork(artType string, artId int64) (err error) {
	var artwork *dtos.ArtworkDTO
	global.JobPool.Submit(func() { artwork, err = apis.GetIllusts(DefaultContext, artId) })
	if err != nil {
		return err
	}

	var row model.PixivArtworkMod
	if err = global.DB().Where(
		model.PixivArtworkMod{Pid: artwork.Pid, ArtType: artType, ArtId: artwork.ArtId},
	).Assign(
		model.PixivArtworkMod{
			Pid:           artwork.Pid,
			ArtId:         artwork.ArtId,
			ArtType:       artType,
			Title:         artwork.Title,
			Description:   artwork.Description,
			PageCount:     artwork.PageCount,
			ViewCount:     artwork.ViewCount,
			LikeCount:     artwork.LikeCount,
			BookmarkCount: artwork.BookmarkCount,
			CreateDate:    &artwork.CreateDate,
		},
	).FirstOrCreate(&row).Error; err != nil {
		return err
	}

	var tagId int64
	for _, tag := range artwork.Tags.Tags {
		global.JobPool.Submit(func() { tagId, err = SyncTag(tag.Jp) })
		if err != nil {
			continue
		}

		if err = markArtworkTag(artwork.Pid, artType, artId, tagId); err != nil {
			continue
		}
	}

	if err = global.DB().Exec(`UPDATE pixiv_account SET art_updated = ? WHERE pid = ? AND art_updated < ?`, artwork.CreateDate.Unix(), artwork.Pid, artwork.CreateDate.Unix()).Error; err != nil {
		return err
	}
	return nil
}

func (s Service) SyncArtWorks(pid int64) (err error) {
	if _, err = s.FlushAccountInfo(pid); err != nil {
		return err
	}

	var profile *dtos.AllProfileDTO
	global.JobPool.Submit(func() { profile, err = apis.GetProfileAll(DefaultContext, pid) })
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
				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork] (%9d):   ERROR,  artType=%6s  artId=%9d  errmsg=%s", pid, artType, artId, err.Error()))
				return err
			}
			global.Logger.Debug().Msg(fmt.Sprintf("[SyncArtWork] (%9d): RUNNING  artType=%6s  artId=%9d  %4d/%4d ", pid, artType, artId, idx+1, len(artIds)))
		}
	}
	return nil
}
