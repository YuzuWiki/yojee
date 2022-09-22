package pixiv_service

import (
	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
)

func (Service) GetArtwork(artType string, artId int64) (*ArtworkDTO, error) {
	// TODO: fix
	artwork := model.PixivArtworkMod{}
	if err := global.DB().Where("art_type=? AND art_id=? AND is_deleted=?", artType, artId, false).First(&artwork).Error; err != nil {
		return nil, err
	}

	tags := make([]struct {
		Jp string `gorm:"column:jp"`
	}, 0)
	if err := global.DB().Raw(`
		SELECT
			artwork.*,
			json_array(
				(
					SELECT
						tag.jp
					FROM pixiv_tag          AS tag
					JOIN pixiv_artwork_tag  AS pat
						ON pat.tag_id = tag.id
					WHERE pat.is_deleted = false
					  AND pat.art_type = artwork.art_type
					  AND pat.art_id = artwork.art_id
				)
			) as tags
		FROM pixiv_artwork AS artwork
		WHERE artwork.art_type=? AND artwork.art_id =?;`,
		global.DATABASE(), global.DATABASE(), artType, artId,
	).Scan(tags).Error; err != nil {
		return nil, err
	}

	retTags := make([]string, 0)
	for i := range tags {
		retTags = append(retTags, tags[i].Jp)
	}

	retArtwork := ArtworkDTO{
		Pid:           artwork.Pid,
		ArtId:         artwork.ArtId,
		ArtType:       artwork.ArtType,
		Title:         artwork.Title,
		Description:   artwork.Description,
		PageCount:     artwork.PageCount,
		ViewCount:     artwork.ViewCount,
		LikeCount:     artwork.ViewCount,
		BookmarkCount: artwork.BookmarkCount,
		CreateDate:    artwork.CreateDate,
		Tags:          retTags,
	}
	return &retArtwork, nil
}
