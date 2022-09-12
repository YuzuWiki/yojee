package model

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
)

type PixivTagMod struct {
	BaseMod

	Jp     string `gorm:"type:VARCHAR(512);column:jp" json:"jp"`
	En     string `gorm:"type:VARCHAR(512);column:en" json:"en"`
	Ko     string `gorm:"type:VARCHAR(512);column:ko" json:"ko"`
	Zh     string `gorm:"type:VARCHAR(512);column:zh" json:"zh"`
	Romaji string `gorm:"type:VARCHAR(512);column:romaji" json:"romaji"`
}

func (PixivTagMod) TableName() string {
	return "pixiv_tag"
}

func (PixivTagMod) FindId(jp string) (int64, error) {
	if len(jp) == 0 {
		return 0, fmt.Errorf("invalid tag")
	}

	var tag struct {
		Id int64 `gorm:"column:id" json:"id"`
	}
	if err := global.DB().Exec(`SELECT id FROM pixiv_tag WHERE jp=? AND is_deleted=0 LIMIT 1;`, jp).Scan(&tag).Error; err != nil {
		return 0, err
	}
	return tag.Id, nil
}

func (PixivTagMod) Insert(jp, romaji, en, ko, zh string) (int64, error) {
	row := &PixivTagMod{
		Jp:     jp,
		En:     en,
		Ko:     ko,
		Zh:     zh,
		Romaji: romaji,
	}

	if err := global.DB().Create(row).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}

type PixivTagTreeMod struct {
	BaseMod

	ParentId int64  `gorm:"type:bigint;column:parent_id" json:"parent_id"`
	ParentJp string `gorm:"type:VARCHAR(512);column:parent_jp" json:"parent_jp"`

	TagId int64  `gorm:"type:bigint;column:tag_id" json:"tag_id"`
	TagJp string `gorm:"type:VARCHAR(512);column:tag_jp" json:"tag_jp"`
}

func (PixivTagTreeMod) TableName() string {
	return "pixiv_tag_tree"
}

func (PixivTagTreeMod) Insert(parentId int64, parentJp string, childId int64, childJp string) (int64, error) {
	row := &PixivTagTreeMod{
		ParentId: parentId,
		ParentJp: parentJp,
		TagId:    childId,
		TagJp:    childJp,
	}
	if err := global.DB().Create(row).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}

type PixivArtworkTagMod struct {
	BaseMod

	ArtId   int64  `gorm:"type:bigint;column:art_id" json:"art_id"`
	ArtType string `gorm:"type:varchar(64);column:art_type" json:"art_type"`

	TagId int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivArtworkTagMod) TableName() string {
	return "pixiv_artwork_tag"
}

func (PixivArtworkTagMod) MarkTag(artType string, artId int64, tagId int64) error {
	row := PixivArtworkTagMod{
		ArtId:   artId,
		ArtType: artType,
		TagId:   tagId,
	}
	if err := global.DB().FirstOrCreate(&row, PixivArtworkTagMod{ArtId: artId, ArtType: artType, TagId: tagId}).Error; err != nil {
		global.Logger.Error().Msg(err.Error())
		return err
	}
	return nil
}
