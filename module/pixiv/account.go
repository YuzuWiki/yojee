package pixiv

import (
	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
)

func (s Service) GetPid(phpSessId string) (pid int64, err error) {
	ctx := Pixivlee.NewContext(phpSessId)

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

func flushAccount(pid int64) (_ *model.PixivAccountMod, err error) {
	data, err := apis.GetAccountInfo(DefaultContext, pid)
	if err != nil {
		return nil, err
	}

	account := &model.PixivAccountMod{}
	if err = global.DB().Where(
		model.PixivAccountMod{Pid: data.UserID, IsDeleted: false},
	).Assign(
		model.PixivAccountMod{
			Pid:       data.UserID,
			Name:      data.Name,
			Avatar:    data.Avatar,
			Region:    data.Region.Name,
			Gender:    data.Gender.Name,
			BirthDay:  data.BirthDay.Name,
			Job:       data.Job.Name,
			Following: data.Following,
			FanboxUrl: "",
		},
	).FirstOrCreate(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func flushFanboxUrl(pid int64) (u string, err error) {
	if u, err = apis.GetFanboxUlr(DefaultContext, pid); err != nil {
		return "", err
	}

	return u, global.DB().Exec(`UPDATE pixiv_account SET fanbox_url = ? WHERE pid = ? AND is_deleted = false`, u, pid).Error
}

func (s Service) FlushAccountInfo(pid int64) (account *model.PixivAccountMod, err error) {
	global.JobPool.Submit(func() { account, err = flushAccount(pid) })
	if err != nil {
		return nil, err
	}

	var fanboxUrl string
	global.JobPool.Submit(func() { fanboxUrl, err = flushFanboxUrl(pid) })
	if err != nil {
		return nil, err
	}
	account.FanboxUrl = fanboxUrl
	return account, nil
}
