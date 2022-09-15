package pixiv_service

import (
	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

func (s Service) GetPid(phpSessId string) (pid int64, err error) {
	ctx := pixiv.NewContext(phpSessId)

	if pid, err = apis.GetAccountPid(ctx); err != nil {
		return 0, err
	}

	return pid, nil
}

func (s Service) GetAccountInfo(pid int64) (_ *model.PixivUserMod, err error) {
	row := &model.PixivUserMod{}
	if err = global.DB().First(row, model.PixivUserMod{PID: pid, IsDeleted: false}).Error; err != nil {
		return nil, err
	}
	return row, nil
}

func (s Service) FlushAccountInfo(pid int64) (_ *model.PixivUserMod, err error) {
	data, err := apis.GetAccountInfo(pixiv.DefaultContext, pid)
	if err != nil {
		return nil, err
	}

	var (
		row = &model.PixivUserMod{
			PID:       data.UserID,
			Name:      data.Name,
			Avatar:    data.Avatar,
			Region:    data.Region.Name,
			Gender:    data.Gender.Name,
			BirthDay:  data.BirthDay.Name,
			Job:       data.Job.Name,
			Following: data.Following,
		}

		db = global.DB()
	)

	if err = db.Table(model.PixivUserMod{}.TableName()).Where("pid = ? AND is_deleted = ?", pid, false).Updates(map[string]interface{}{"is_deleted": true}).Error; err != nil {
		return nil, err
	}

	if err = db.Create(row).Error; err != nil {
		return nil, err
	}
	return row, err
}
