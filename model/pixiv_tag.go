package model

type PixivTagMod struct {
	BaseMod

	Name   string `gorm:"type:VARCHAR(512);column:name" json:"name"`
	Romaji string `gorm:"type:VARCHAR(512);column:romaji" json:"romaji"`
}

func (PixivTagMod) TableName() string {
	return "pixiv_tag"
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
