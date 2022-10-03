package pixiv_service

import (
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
)

func (Service) GetFollowing(pid int64, limit, offset int) (_ *[]model.PixivAccountMod, err error) {
	follows := make([]model.PixivAccountMod, 0)
	global.Logger.Debug().Msg(fmt.Sprintf("[GetFollowing] (%9d): limit=%2d offset=%2d", pid, limit, offset))
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

func syncFollowing(pid int64, followingPid int64) (err error) {
	var account *model.PixivAccountMod
	if account, err = flushAccount(followingPid); err != nil {
		return err
	}

	if err = global.DB().Where("pid = ? AND followed_pid = ?", pid, account.Pid).FirstOrCreate(&model.PixivFollowMod{PID: pid, FollowedPid: followingPid, IsDeleted: false}).Error; err != nil {
		return err
	}
	return nil
}

func (s Service) SyncFollowing(pid int64) (_ int, err error) {
	var (
		limit  = 50
		offset = 0
		total  = offset + 1
		body   *dtos.FollowingDTO
	)
	global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%9d): BEGIN", pid))

	// un-mark following status
	if err = global.DB().Exec("UPDATE pixiv_follow SET is_deleted = true WHERE pid = ?;", pid).Error; err != nil {
		return 0, err
	}

	// re-mark following status
	for (offset) <= total {
		if body, err = apis.GetFollowing(pixiv.DefaultContext, pid, limit, offset); err != nil {
			global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%9d): ERROR, GetFollowing errmsg=%s", pid, err.Error()))
			return 0, err
		}
		_offset := offset

		total = body.Total
		offset += limit

		for idx, u := range body.Users {
			global.JobPool.Submit(func() {
				if err = syncFollowing(pid, u.UserID); err != nil {
					global.Logger.Error().Err(err).Msg(fmt.Sprintf("[SyncFollowing] (%9d): Error  following_pid=%d  %3d/%3d, FlushAccountInfo errmsg=%s", pid, u.UserID, idx+_offset+1, total, err.Error()))
				} else {
					global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%9d): Doing  following_pid=%9d  %3d/%3d", pid, u.UserID, idx+_offset+1, total))
				}
			})

			global.JobPool.Submit(func() { _, _ = flushFanboxUrl(pid) })
		}
	}
	global.Logger.Debug().Msg(fmt.Sprintf("[SyncFollowing] (%9d): DONE, total=%d", pid, total))
	return total, nil
}
