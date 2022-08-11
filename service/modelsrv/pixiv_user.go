package modelsrv

import (
	"github.com/jinzhu/gorm"

	"github.com/YuzuWiki/yojee/global"
	"github.com/YuzuWiki/yojee/model"
	"github.com/YuzuWiki/yojee/module/pixiv/apis"
)

type PixivUser struct{}

// FindUser return data from database
func (srv *PixivUser) findUser(db *gorm.DB, pid int64) (*model.PixivUserMod, error) {

	var user model.PixivUserMod
	if err := db.Raw(
		"SELECT * FROM pixiv_user WHERE pid=? AND is_deleted=false LIMIT 1;", pid,
	).Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (srv *PixivUser) FindUser(pid int64) (*model.PixivUserMod, error) {
	return srv.findUser(global.DB(), pid)
}

// InsertUser will insert user data
func (srv *PixivUser) insertUser(tx *gorm.DB, info apis.UserInfoDTO) error {
	// 删除原有记录
	if err := tx.Raw("UPDATE pixiv_user SET is_deleted=true WHERE pid=? AND is_deleted=false LIMIT 1;").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&model.PixivUserMod{
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

func (srv *PixivUser) InsertUser(info apis.UserInfoDTO) error {
	db := global.DB()

	user, err := srv.findUser(db, info.UserID)
	if err != nil {
		return err
	}

	if user.Name == info.Name &&
		user.Avatar == info.Avatar &&
		user.Gender == info.Gender.Name &&
		user.Region == info.Region.Name &&
		user.Following == info.Following {
		// 若无更新， 则pass
		return nil
	}

	return srv.insertUser(db.Begin(), info)
}
