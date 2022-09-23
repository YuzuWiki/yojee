package pixiv_service

import (
	"fmt"
	"strconv"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
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
				ON pat.tag_id    = tag.id
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

func (s Service) SyncArtWork(pid int64) error {
	if _, err := s.FlushAccountInfo(pid); err != nil {
		return err
	}

	global.JobPool.Submit(func() {
		profile, err := apis.GetProfileAll(pixiv.DefaultContext, pid)
		if err != nil {
			global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork]  ERROR: (%d) GetProfileAll error, %s", pid, err.Error()))
			return
		}

		for illustId := range profile.Illusts {
			var artId int64
			if artId, err = strconv.ParseInt(illustId, 10, 64); err != nil {
				global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork]  ERROR: (%d)-(%s)  GetIllusts error, %s", pid, illustId, err.Error()))
				continue
			}

			global.JobPool.Submit(func() {
				data, err := apis.GetIllusts(pixiv.DefaultContext, artId)
				if err != nil {
					global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork]  ERROR: (%d)-(%s)  GetIllusts error, %s", pid, illustId, err.Error()))
					return
				}

				if _, err = (model.PixivArtworkMod{}).Insert(*data); err != nil {
					global.Logger.Error().Msg(fmt.Sprintf("[SyncArtWork]  ERROR: (%d)-(%s)  GetIllusts error, %s", pid, illustId, err.Error()))
					return
				}

				for idx := range data.Tags.Tags {
					tag := data.Tags.Tags[idx]

					tagId, err := model.PixivTagMod{}.GetId(tag.Jp)
					if err != nil {
						// TODO tag info
					}
					global.Logger.Error().Msg(fmt.Sprintf("%d", tagId))
				}
			})
		}
	})
	return nil
}
