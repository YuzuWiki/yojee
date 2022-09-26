package pixiv_service

import (
	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func findTagId(jp string) (tagId int64, err error) {
	if err = global.DB().Exec(`SELECT tag_id FROM pixiv_tag WHERE jp=? AND is_deleted=false LIMIT 1;`, jp).Find(&tagId).Error; err != nil {
		return 0, err
	}
	return tagId, nil
}

func syncTag(jp string) (tagId int64, err error) {
	var tag *dtos.TagDTO
	if tag, err = apis.GetTag(pixiv.DefaultContext, jp); err != nil {
		return 0, err
	}

	row := model.PixivTagMod{
		TagId:     tag.Digest.Id,
		Jp:        tag.Jp,
		En:        tag.Translation[tag.Jp].En,
		Ko:        tag.Translation[tag.Jp].Ko,
		Zh:        tag.Translation[tag.Jp].Zh,
		Romaji:    tag.Translation[tag.Jp].Romaji,
		IsDeleted: false,
	}
	if err = global.DB().Where(model.PixivTagMod{TagId: tag.Digest.Id}).FirstOrCreate(&row).Error; err != nil {
		return 0, err
	}
	return row.TagId, nil
}

func SyncTag(jp string) (tagId int64, err error) {
	if tagId, err = findTagId(jp); err == nil {
		return tagId, nil
	}

	if tagId, err = syncTag(jp); err != nil {
		return 0, err
	}
	return tagId, nil
}

func markArtworkTag(pid int64, artType string, artID int64, tagId int64) error {
	row := model.PixivArtworkTagMod{
		Pid:     pid,
		ArtId:   artID,
		ArtType: artType,
		TagId:   tagId,
	}
	if err := global.DB().FirstOrCreate(&row, row).Error; err != nil {
		return err
	}
	return nil
}
