package modelsrv

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivTags struct{}

func (PixivTags) FindByArtId(category string, artId int64) (*[]model.PixivTagMod, error) {
	if category != apis.Illust && category != apis.Manga && category != apis.Novel {
		return nil, fmt.Errorf("category(%s) not support", category)
	}

	db := global.DB()

	var tags []model.PixivTagMod
	if err := db.Exec(`
		SELECT
			tag.id          AS id,
			tag.name        AS name,
			tag.created_at  AS created_at,
			tag.updated_at  AS updated_at,
			tag.is_deleted  AS is_deleted
		FROM pixiv_tag 			AS tag
		JOIN pixiv_artwork_tag  AS pag
			ON tag.id=pag.tag_id AND pag.is_deleted=false
		WHERE pag.art_type=? AND pag.art_id=?;`, category, artId,
	).Scan(&tags).Error; err != nil {
		return nil, err
	}

	return &tags, nil
}

func (PixivTags) FindByName(name string) (*[]model.PixivTagMod, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("miss tag name")
	}

	db := global.DB()

	var tags []model.PixivTagMod
	if err := db.Exec(`SELECT * FROM pixiv_tag WHERE name=? AND is_deleted=false;`, name).Scan(&tags).Error; err != nil {
		return nil, err
	}

	return &tags, nil
}

func (PixivTags) InsertTags(names ...string) int {
	var cnt int
	db := global.DB()
	for _, name := range names {
		if err := db.Create(model.PixivTagMod{Name: name}).Error; err != nil {
			global.Logger.Warn().Msg(fmt.Sprintf("insert (%s) error,  %s", name, err.Error()))
		} else {
			cnt++
		}
	}
	return cnt
}
