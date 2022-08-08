package model

type PixivTagMod struct {
	BaseMod

	Name string `gorm:"column:name" json:"name"`
}

func (PixivTagMod) TableName() string {
	return "pixiv_tag"
}

type PixivIllustTagMod struct {
	BaseMod

	IllustId int64 `gorm:"type:bigint;column:illust_id" json:"illust_id"`
	TagId    int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivIllustTagMod) TableName() string {
	return "pixiv_illust_tag"
}

type PixivMangaTagMod struct {
	BaseMod

	MangaId int64 `gorm:"type:bigint;column:manga_id" json:"manga_id"`
	TagId   int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivMangaTagMod) TableName() string {
	return "pixiv_manga_tag"
}

type PixivNovelTagMod struct {
	BaseMod

	NovelId int64 `gorm:"type:bigint;column:novel_id" json:"novel_id"`
	TagId   int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}

func (PixivNovelTagMod) TableName() string {
	return "pixiv_novel_tag"
}
