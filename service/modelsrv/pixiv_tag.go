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

func (PixivTag) FindID(name string) (int64, error) {
	if len(name) == 0 {
		return 0, fmt.Errorf("miss tag name")
	}

	if tagId, err := global.RDB().Get(_CacheTagPrefix + name).Result(); err == nil {
		return strconv.ParseInt(tagId, 10, 0)
	}

	var tag model.PixivTagMod
	if err := global.DB().Exec(`SELECT * FROM pixiv_tag WHERE name=? AND is_deleted=false LIMIT 1;`, name).Scan(&tag).Error; err != nil {
		return 0, err
	}

	global.RDB().Set(_CacheTagPrefix+tag.Name, tag.ID, 5*60*time.Second)
	return int64(tag.ID), nil
}

func (PixivTag) Insert(name, romaji string) (int64, error) {
	rdb := global.RDB()
	if tagId, err := rdb.Get(_CacheTagPrefix + name).Result(); err == nil {
		return strconv.ParseInt(tagId, 10, 0)
	}

	row := model.PixivTagMod{Name: name}
	if err := global.DB().FirstOrCreate(&row, &model.PixivTagMod{Name: name}).Error; err != nil {
		return 0, err
	}

	// set 缓存
	rdb.Set(_CacheTagPrefix+row.Name, row.ID, 5*60*time.Second)
	return int64(row.ID), nil
}
