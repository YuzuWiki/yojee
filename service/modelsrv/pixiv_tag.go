package modelsrv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
)

const _CacheTagPrefix = "pixiv:tag:cache:"

type PixivTag struct{}

func (PixivTag) FindID(jp string) (int64, error) {
	if len(jp) == 0 {
		return 0, fmt.Errorf("miss tag jp")
	}

	if tagId, err := global.RDB().Get(_CacheTagPrefix + jp).Result(); err == nil {
		return strconv.ParseInt(tagId, 10, 0)
	}

	var tag model.PixivTagMod
	if err := global.DB().Exec(`SELECT * FROM pixiv_tag WHERE jp=? AND is_deleted=0 LIMIT 1;`, jp).Scan(&tag).Error; err != nil {
		return 0, err
	}

	global.RDB().Set(_CacheTagPrefix+tag.Jp, tag.ID, 5*60*time.Second)
	return int64(tag.ID), nil
}

func (PixivTag) Insert(name, romaji string) (int64, error) {
	rdb, db := global.RDB(), global.DB()
	if tagId, err := rdb.Get(_CacheTagPrefix + name).Result(); err == nil {
		return strconv.ParseInt(tagId, 10, 0)
	}

	row := model.PixivTagMod{Jp: name, Romaji: romaji}
	err := db.Create(&row).Error
	if err != nil && !global.IsDuplicateEntry(err) {
		return 0, err
	}

	if err := db.First(&row, &model.PixivTagMod{Jp: name}).Error; err != nil {
		return 0, err
	}

	// set 缓存
	rdb.Set(_CacheTagPrefix+row.Jp, row.ID, 5*60*time.Second)
	return int64(row.ID), nil
}
