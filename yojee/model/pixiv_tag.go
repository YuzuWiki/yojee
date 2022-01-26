package model

type PixivTagMod struct {
	BaseMod

	Name string `gorm:"column:name" json:"name"`
}

type PixivIllustTagMod struct {
	BaseMod

	IllustId int64 `gorm:"type:bigint;column:illust_id" json:"illust_id"`
	TagId    int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}
type PixivMangaTagMod struct {
	BaseMod

	MangaId int64 `gorm:"type:bigint;column:manga_id" json:"manga_id"`
	TagId   int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}
type PixivNovelTagMod struct {
	BaseMod

	NovelId int64 `gorm:"type:bigint;column:novel_id" json:"novel_id"`
	TagId   int64 `gorm:"type:bigint;column:tag_id" json:"tag_id"`
}
