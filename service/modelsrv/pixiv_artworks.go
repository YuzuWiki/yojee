package modelsrv

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivArtworks struct{}

func find(pid int64, artType string) (*[]model.PixivArtworkMod, error) {
	db := global.DB()

	var artworks []model.PixivArtworkMod
	if err := db.Exec("SELECT * FROM pixiv_artwork WHERE pid=? AND art_type=? AND is_deleted=false;", pid, artType).Find(&artworks).Error; err != nil {
		return nil, err
	}
	return &artworks, nil
}
func (PixivArtworks) FindIllustrates(pid int64) (*[]model.PixivArtworkMod, error) {
	return find(pid, apis.Illust)
}

func (PixivArtworks) FindMangas(pid int64) (*[]model.PixivArtworkMod, error) {
	return find(pid, apis.Manga)
}

func (PixivArtworks) FindNovels(pid int64) (*[]model.PixivArtworkMod, error) {
	return find(pid, apis.Novel)
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

func insert(artType string, data apis.ArtworkDTO) (int64, error) {
	row := model.PixivArtworkMod{
		Pid:           data.UserId,
		ArtId:         data.Id,
		ArtType:       artType,
		Title:         data.Title,
		Description:   data.Description,
		ViewCount:     data.ViewCount,
		LikeCount:     data.LikeCount,
		BookmarkCount: data.BookmarkCount,
		CreateDate:    &data.CreateDate,
	}
	if err := global.DB().FirstOrCreate(&row, model.PixivArtworkMod{Pid: data.UserId, ArtType: artType, ArtId: data.Id}).Error; err != nil {
		global.Logger.Error().Msg(fmt.Sprintf("insert illust(%d) error,  %s", data.Id, err.Error()))
		return 0, err
	}
	return int64(row.ID), nil
}

func (PixivArtworks) InsertIllust(data apis.ArtworkDTO) (int64, error) {
	return insert(apis.Illust, data)
}

func (PixivArtworks) InsertManga(data apis.ArtworkDTO) (int64, error) {
	return insert(apis.Manga, data)
}

func (PixivArtworks) InsertNovel(data apis.ArtworkDTO) (int64, error) {
	return insert(apis.Novel, data)
}
