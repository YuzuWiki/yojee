package model

import (
	"strings"
	"time"

	"github.com/YuzuWiki/yojee/global"
)

const _TagIdKey = "pixiv:tag:id:"

type PixivTagMod struct {
	ID        uint64     `gorm:"type:timestamp;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	Jp     string `gorm:"type:VARCHAR(512);column:jp" json:"jp"`
	En     string `gorm:"type:VARCHAR(512);column:en" json:"en"`
	Ko     string `gorm:"type:VARCHAR(512);column:ko" json:"ko"`
	Zh     string `gorm:"type:VARCHAR(512);column:zh" json:"zh"`
	Romaji string `gorm:"type:VARCHAR(512);column:romaji" json:"romaji"`
}

func (PixivTagMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_tag"}, ".")
}

func (PixivTagMod) GetId(jp string) (tagId int64, err error) {
	if err = global.DB().Exec(`SELECT id FROM pixiv_tag WHERE jp=? AND is_deleted=false LIMIT 1;`, jp).Find(&tagId).Error; err != nil {
		return 0, err
	}
	return tagId, nil
}

func (PixivTagMod) Find(artType string, artId int64) (tags *[]PixivTagMod, err error) {
	if err = global.DB().Exec(`
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
		JOIN pixiv_artwork_tag  	AS pag
			ON tag.id=pag.tag_id AND pag.is_deleted=false
		WHERE pag.art_type=? AND pag.art_id=?;`, artType, artId,
	).Scan(tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (PixivTagMod) Insert(jp, romaji, en, ko, zh string) (int64, error) {
	tag := &PixivTagMod{
		Jp:     jp,
		En:     en,
		Ko:     ko,
		Zh:     zh,
		Romaji: romaji,
	}

	if err := global.DB().FirstOrCreate(tag, &PixivTagMod{Jp: jp, IsDeleted: false}).Error; err != nil {
		return 0, err
	}
	return int64(tag.ID), nil
}

type PixivTagTreeMod struct {
	ID        uint64     `gorm:"type:timestamp;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	ParentId int64  `gorm:"type:bigint;column:parent_id" json:"parent_id"`
	ParentJp string `gorm:"type:VARCHAR(512);column:parent_jp" json:"parent_jp"`

	TagId int64  `gorm:"type:bigint;column:tag_id" json:"tag_id"`
	TagJp string `gorm:"type:VARCHAR(512);column:tag_jp" json:"tag_jp"`
}

func (PixivTagTreeMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_tag_tree"}, ".")
}

func (PixivTagTreeMod) Insert(parentId int64, parentJp string, childId int64, childJp string) (int64, error) {
	row := &PixivTagTreeMod{
		ParentId: parentId,
		ParentJp: parentJp,
		TagId:    childId,
		TagJp:    childJp,
	}
	if err := global.DB().FirstOrCreate(row, &PixivTagTreeMod{ParentId: parentId, TagId: childId, IsDeleted: false}).Error; err != nil {
		return 0, err
	}
	return int64(row.ID), nil
}

type PixivArtworkTagMod struct {
	ID        uint64     `gorm:"type:timestamp;primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:milli;column:updated_at"  json:"updated_at"`
	IsDeleted bool       `gorm:"type:bool;default:false;column:is_deleted" json:"is_deleted"`

	ArtId   int64  `gorm:"type:bigint;column:art_id" json:"art_id"`
	ArtType string `gorm:"type:varchar(64);column:art_type" json:"art_type"`

	TagId int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivArtworkTagMod) TableName() string {
	return strings.Join([]string{global.DATABASE(), "pixiv_artwork_tag"}, ".")
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
