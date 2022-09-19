package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

func (Service) GetFollowing(pid int64, limit, offset int) (_ *[]model.PixivAccountMod, err error) {
	follows := make([]model.PixivAccountMod, 0)
	global.Logger.Debug().Msg(fmt.Sprintf("[DEBUG] pid=%d limit=%d offset=%d", pid, limit, offset))
	if err = global.DB().Raw(`
		SELECT
			account.*
		FROM pixiv_account AS account
		JOIN pixiv_follow  AS follow
			ON follow.followed_pid = account.pid 
			   AND follow.pid = ?
			   AND follow.is_deleted = false
		WHERE account.is_deleted = false
		LIMIT ? OFFSET ?;`, pid, limit, offset,
	).Scan(&follows).Error; err != nil {
		return nil, err
	}

	return &follows, nil
}

func (s Service) SyncFollowing(pid int64) (int, error) {
	global.Logger.Debug().Msg(fmt.Sprintf("[DEBUG] SyncFollowing: (%d) BEGIN", pid))
	var (
		limit  = 50
		offset = 0
		total  = offset + 1
	)

	// un-mark following status
	if err := global.DB().Exec("UPDATE pixiv_follow SET is_deleted = true WHERE pid = ?;", pid).Error; err != nil {
		return 0, err
	}

	// re-mark following status
	for (offset) <= total {
		body, err := apis.GetFollowing(pixiv.DefaultContext, pid, limit, offset)
		if err != nil {
			global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%d): ERROR, GetFollowing %s", pid, err.Error()))
			return 0, err
		}

		total = body.Total
		offset += limit

		for _, u := range body.Users {
			global.JobPool.Submit(func() {
				var (
					account *model.PixivAccountMod
					e       error
				)
				if account, e = s.FlushAccountInfo(u.UserID); e != nil {
					global.Logger.Error().Err(e).Msg(fmt.Sprintf("[SyncFollowing] (%d): ERROR, %s", u.UserID, e.Error()))
					return
				}

				if e = global.DB().Where("pid = ? AND followed_pid = ?", pid, account.Pid, false).FirstOrCreate(&model.PixivFollowMod{PID: pid, FollowedPid: u.UserID, IsDeleted: false}).Error; e != nil {
					global.Logger.Error().Err(e).Msg(fmt.Sprintf("[SyncFollowing] (%d): ERROR, mark following Fail, pid=%d followed_pid=%d", u.UserID, pid, u.UserID))
					return
				}
				global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%d): DONE, relation synced", u.UserID))
			})
		}
	}
	global.Logger.Debug().Msg(fmt.Sprintf("[DEBUG] SyncFollowing: (%d) DONE, total = %d", pid, total))
	return total, nil
}
