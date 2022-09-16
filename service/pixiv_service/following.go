package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

func (Service) GetFollowing(pid int64, limit, offset int) (follows *[]model.PixivAccountMod, err error) {
	if err = global.DB().Exec(`
	SELECT
		account.*
	FROM pixiv_account AS account
	JOIN pixiv_follow  AS follow
		ON follow.followed_pid = account.pid 
		   AND follow.pid = ?
		   AND follow.is_deleted = false
	WHERE account.is_deleted = false
	LIMIT ? OFFSET ?;`, pid, limit, offset,
	).Scan(follows).Error; err != nil {
		return nil, err
	}

	return follows, nil
}

func (s Service) SyncFollowing(pid int64) (total int, err error) {
	var (
		limit  = 50
		offset = 0
	)

	total = limit + 1
	for (limit + offset) <= total {
		body, err := apis.GetFollowing(pixiv.DefaultContext, pid, 50, 0)
		if err != nil {
			return 0, err
		}

		total = body.Total
		offset += limit

		for _, u := range body.Users {
			global.JobPool.Submit(func() {
				global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%d): BEGIN, relation syncing...", u.UserID))
				var (
					account *model.PixivAccountMod
					e       error
				)
				if account, e = s.FlushAccountInfo(u.UserID); e != nil {
					global.Logger.Error().Err(e).Msg(fmt.Sprintf("[SyncFollowing] (%d): ERROR, sync account Fail", u.UserID))
					return
				}

				if e = global.DB().Where("pid = ? AND followed_pid = ? AND is_deleted = ?", pid, account.Pid, false).FirstOrCreate(&model.PixivFollowMod{PID: pid, FollowedPid: u.UserID}).Error; e != nil {
					global.Logger.Error().Err(e).Msg(fmt.Sprintf("[SyncFollowing] (%d): ERROR, mark following Fail, pid=%d followed_pid=%d", u.UserID, pid, u.UserID))
					return
				}
				global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%d): DONE, relation synced", u.UserID))
			})
		}
	}

	return total, nil
}
