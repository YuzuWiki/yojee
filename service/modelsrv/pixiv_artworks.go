package modelsrv

import (
	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
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
