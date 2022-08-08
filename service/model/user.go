package model

import (
	"errors"
	"fmt"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivUser struct{}

// FindUser return data from database
func (srv *PixivUser) FindUser(pid int64) (*model.PixivUserMod, error) {
	db := global.DB()

	var user model.PixivUserMod
	if db.Raw(
		"SELECT * FROM yojee.pixiv_user WHERE pid = ? AND is_deleted=0 LIMIT 1;", pid,
	).Scan(&user).RecordNotFound() {
		return &user, nil
	}
	return nil, errors.New(fmt.Sprintf("not found user, pid(%d)", pid))
}

// InsertUser will insert user data
func (srv *PixivUser) InsertUser(info apis.UserInfoDTO, extra apis.ExtraDTO) error {
	tx := global.DB().Begin()

	// 删除原有记录
	if err := tx.Exec("UPDATE pixiv_user SET is_deleted = 1 WHERE pid=? AND is_deleted=0 LIMIT 1;").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&model.PixivUserMod{
		PID:       info.UserID,
		Name:      info.Name,
		Avatar:    info.Avatar,
		Region:    info.Region.Region,
		Gender:    info.Gender.Name,
		Following: extra.Following,
		Followers: extra.Followers,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
