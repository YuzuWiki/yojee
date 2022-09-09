package modelsrv

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

type PixivArtwork struct{}

func (PixivArtwork) Find(artType string, pid int64) (*[]model.PixivArtworkMod, error) {
	var artworks []model.PixivArtworkMod
	if err := global.DB().Exec("SELECT * FROM pixiv_artwork WHERE pid=? AND art_type=? AND is_deleted=false;", pid, artType).Find(&artworks).Error; err != nil {
		return nil, err
	}
	return &artworks, nil
}

func (PixivArtwork) FindTags(artType string, artId int64) (*[]model.PixivTagMod, error) {
	var tags []model.PixivTagMod
	if err := global.DB().Exec(`
		SELECT
			tag.id          AS id,
			tag.jp          AS jp,
			tag.en          AS en,
			tag.ko          AS ko,
			tag.zh          AS zh,
			tag.created_at  AS created_at,
			tag.updated_at  AS updated_at,
			tag.is_deleted  AS is_deleted
		FROM pixiv_tag 			AS tag
		JOIN pixiv_artwork_tag  AS pag
			ON tag.id=pag.tag_id AND pag.is_deleted=false
		WHERE pag.art_type=? AND pag.art_id=?;`, artType, artId,
	).Scan(&tags).Error; err != nil {
		return nil, err
	}

	return &tags, nil
}

func (PixivArtwork) Insert(data dtos.ArtworkDTO) (int64, error) {
	row := model.PixivArtworkMod{
		Pid:           data.Pid,
		ArtId:         data.ArtId,
		ArtType:       data.ArtType,
		Title:         data.Title,
		Description:   data.Description,
		ViewCount:     data.ViewCount,
		LikeCount:     data.LikeCount,
		BookmarkCount: data.BookmarkCount,
		CreateDate:    &data.CreateDate,
	}
	if err := global.DB().FirstOrCreate(&row, model.PixivArtworkMod{Pid: data.Pid, ArtType: data.ArtType, ArtId: data.ArtId}).Error; err != nil {
		global.Logger.Error().Msg(fmt.Sprintf("insert illust(%d) error,  %s", data.ArtId, err.Error()))
		return 0, err
	}
	return int64(row.ID), nil
}

func (PixivArtwork) MarkTag(artType string, artId int64, tagId int64) error {
	row := model.PixivArtworkTagMod{
		ArtId:   artId,
		ArtType: artType,
		TagId:   tagId,
	}
	if err := global.DB().FirstOrCreate(&row, model.PixivArtworkTagMod{ArtId: artId, ArtType: artType, TagId: tagId}).Error; err != nil {
		global.Logger.Error().Msg(err.Error())
		return err
	}
	return nil
}
