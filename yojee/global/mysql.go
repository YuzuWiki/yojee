package global

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	if DB == nil {
		db, err := gorm.Open("mysql", "root:2125031jun@(192.168.123.122:3306)/yojee?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("MySQL: init db fail ...")
		}
		DB = db
	}
	return DB
}

func CloseDB() error {
	if DB == nil {
		return nil
	}

	if err := DB.Close(); err != nil {
		panic("MySQL: close db fail ...")
	}

	return nil
}
