package modelsrv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
)

const _CacheTagPrefix = "pixiv:tag:cache:"

type PixivTags struct{}


func (PixivTags) FindID(name string) (uint64, error) {
	if len(name) == 0 {
		return 0, fmt.Errorf("miss tag name")
	}

	if tagId, err := global.RDB().Get(_CacheTagPrefix + name).Result(); err == nil {
		return strconv.ParseUint(tagId, 10, 0)
	}

	var tag model.PixivTagMod
	if err := global.DB().Exec(`SELECT * FROM pixiv_tag WHERE name=? AND is_deleted=false LIMIT 1;`, name).Scan(&tag).Error; err != nil {
		return 0, err
	}

	global.RDB().Set(_CacheTagPrefix+tag.Name, tag.ID, 5*60*time.Second)
	return tag.ID, nil
}

func (PixivTags) Insert(names ...string) int {
	var cnt int
	db, rdb := global.DB(), global.RDB()
	for _, name := range names {
		row := model.PixivTagMod{Name: name}
		if err := db.Create(&row).Error; err != nil {
			global.Logger.Warn().Msg(fmt.Sprintf("insert (%s) error,  %s", name, err.Error()))
		} else {
			cnt++
		}

		// set 缓存
		rdb.Set(_CacheTagPrefix+row.Name, row.ID, 5*60*time.Second)
	}
	return cnt
}
