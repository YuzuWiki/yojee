package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
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
		)as tags
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
