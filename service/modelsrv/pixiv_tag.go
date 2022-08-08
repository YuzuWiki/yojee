package modelsrv

type PixivTags struct{}

//
//func (PixivTags) FindTags(category string, id int64) (*[]model.PixivTagMod, error) {
//	db := global.DB()
//
//	sql := `
//	SELECT
//		tag.id          AS id,
//		tag.name        AS name,
//		tag.created_at  AS created_at,
//		tag.updated_at  AS updated_at,
//		tag.is_deleted  AS is_deleted
//	FROM pixiv_tag AS tag
//	JOIN pixiv_artwork_tag  AS pag
//		ON tag.id=pag.tag_id AND pag.is_deleted=false
//	WHERE pag.art_type=? AND pag.art_id=?;
//	`
//
//	var (
//		tableName string
//		tid       string
//	)
//	switch category {
//	case apis.Illust:
//		tableName, tid = model.PixivIllustTagMod{}.TableName(), "illust_id"
//	case apis.Manga:
//		tableName, tid = model.PixivMangaTagMod{}.TableName(), "manga_id"
//	case apis.Novel:
//		tableName, tid = model.PixivNovelTagMod{}.TableName(), "novel_id"
//
//	default:
//		return nil, errors.New(fmt.Sprintf("category(%s) not support", category))
//	}
//
//}
