package model

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

type PixivTagTaxonMod struct {
	BaseMod

	PerId int64  `gorm:"type:bigint;column:per_id" json:"per_id"`
	PerJp string `gorm:"type:VARCHAR(512);column:per_jp" json:"per_jp"`

	TagId int64  `gorm:"type:bigint;column:tag_id" json:"tag_id"`
	TagJp string `gorm:"type:VARCHAR(512);column:tag_jp" json:"tag_jp"`
}

func (PixivTagTaxonMod) TableName() string {
	return "pixiv_tag_taxon"
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
