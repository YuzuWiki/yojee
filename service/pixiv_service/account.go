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

func (s Service) GetAccountInfo(pid int64) (_ *model.PixivAccountMod, err error) {
	row := &model.PixivAccountMod{}
	if err = global.DB().First(row, model.PixivAccountMod{Pid: pid, IsDeleted: false}).Error; err != nil {
		return nil, err
	}
	return row, nil
}

func (s Service) FlushAccountInfo(pid int64) (_ *model.PixivAccountMod, err error) {
	data, err := apis.GetAccountInfo(pixiv.DefaultContext, pid)
	if err != nil {
		return nil, err
	}

	account := &model.PixivAccountMod{}
	if err = global.DB().Where(
		model.PixivAccountMod{Pid: data.UserID, IsDeleted: false},
	).Assign(
		// TODO: 可能有问题
		model.PixivAccountMod{
			Pid:       data.UserID,
			Name:      data.Name,
			Avatar:    data.Avatar,
			Region:    data.Region.Name,
			Gender:    data.Gender.Name,
			BirthDay:  data.BirthDay.Name,
			Job:       data.Job.Name,
			Following: data.Following,
		},
	).FirstOrCreate(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
