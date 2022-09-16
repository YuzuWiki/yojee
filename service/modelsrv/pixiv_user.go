package modelsrv

import (
	"github.com/YuzuWiki/yojee/module/pixiv/dtos"
	"gorm.io/gorm"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
)

type PixivUser struct{}

// FindUser return data from database
func (srv *PixivUser) findUser(db *gorm.DB, pid int64) (*model.PixivAccountMod, error) {
	var user model.PixivAccountMod
	if err := db.Raw(
		"SELECT * FROM pixiv_account WHERE pid=? AND is_deleted=false LIMIT 1;", pid,
	).Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (srv *PixivUser) FindUser(pid int64) (*model.PixivAccountMod, error) {
	return srv.findUser(global.DB(), pid)
}

// InsertUser will insert user data
func (srv *PixivUser) insertUser(tx *gorm.DB, info dtos.UserInfoDTO) error {
	// 删除原有记录
	if err := tx.Exec("UPDATE pixiv_account SET is_deleted=true WHERE pid=? AND is_deleted=false LIMIT 1;", info.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&model.PixivAccountMod{
		PID:       info.UserID,
		Name:      info.Name,
		Avatar:    info.Avatar,
		Region:    info.Region.Name,
		Gender:    info.Gender.Name,
		Following: info.Following,
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

func (srv *PixivUser) InsertUser(info dtos.UserInfoDTO) error {
	tx := global.DB().Begin()

	user, _ := srv.findUser(tx, info.UserID)
	if user != nil && user.Name == info.Name &&
		user.Avatar == info.Avatar &&
		user.Gender == info.Gender.Name &&
		user.Region == info.Region.Name &&
		user.Following == info.Following {
		// 若无更新， 则pass
		return nil
	}

	return srv.insertUser(tx, info)
}
