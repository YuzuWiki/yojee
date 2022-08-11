package modelsrv

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivArtworks struct{}

func (PixivArtworks) FindIllustrates(pid int64) (*[]model.PixivIllustMod, error) {
	db := global.DB()

	var illustrates []model.PixivIllustMod
	if err := db.Exec("SELECT * FROM pixiv_illust WHERE is_deleted=false AND pid=?;", pid, false).Find(&illustrates).Error; err != nil {
		return nil, err
	}
	return &illustrates, nil
}

func (PixivArtworks) FindMangas(pid int64) (*[]model.PixivMangaMod, error) {
	db := global.DB()

	var mangas []model.PixivMangaMod
	if err := db.Exec("SELECT * FROM pixiv_manga WHERE is_deleted=false AND pid=?;", pid, false).Find(&mangas).Error; err != nil {
		return nil, err
	}
	return &mangas, nil
}

func (PixivArtworks) FindNovels(pid int64) (*[]model.PixivNovelMod, error) {
	db := global.DB()

	var mangas []model.PixivNovelMod
	if err := db.Exec("SELECT * FROM pixiv_novel WHERE is_deleted=false AND pid=?;", pid, false).Find(&mangas).Error; err != nil {
		return nil, err
	}
	return &mangas, nil
}

func findTags(category string, artId int64) (*[]model.PixivTagMod, error) {
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

func (PixivArtworks) FindIllustTags(artId int64) (*[]model.PixivTagMod, error) {
	return findTags(apis.Illust, artId)
}

func (PixivArtworks) FindMangaTags(artId int64) (*[]model.PixivTagMod, error) {
	return findTags(apis.Manga, artId)
}

func (PixivArtworks) FindNovelTags(artId int64) (*[]model.PixivTagMod, error) {
	return findTags(apis.Novel, artId)
}

func (PixivArtworks) InsertIllust(data apis.ArtworkIllustDTO) (int64, error) {
	row := model.PixivIllustMod{
		Pid:           data.UserId,
		IllustId:      data.Id,
		Title:         data.Title,
		Description:   data.Description,
		ViewCount:     data.ViewCount,
		LikeCount:     data.LikeCount,
		BookmarkCount: data.BookmarkCount,
		CreateDate:    &data.CreateDate,
	}
	if err := global.DB().FirstOrCreate(&row, model.PixivIllustMod{Pid: data.UserId, IllustId: data.Id}).Error; err != nil {
		global.Logger.Error().Msg(fmt.Sprintf("insert illust(%d) error,  %s", data.Id, err.Error()))
		return 0, err
	}
	return int64(row.ID), nil
}

func (PixivArtworks) InsertManga() {
	//	TODO
}

func (PixivArtworks) InsertNovel() {
	// TODO
}
