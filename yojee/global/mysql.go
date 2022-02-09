package global

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dB *gorm.DB

func InitDB() *gorm.DB {
	if dB == nil {
		db, err := gorm.Open("mysql", "root:2125031jun@(192.168.123.122:3306)/yojee?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("MySQL: init db fail ...")
		}
		dB = db
	}
	return dB
}

func NewDB() *sql.DB  {
	return dB.DB()
}

func CloseDB() error {
	if dB == nil {
		return nil
	}

	if err := dB.Close(); err != nil {
		panic("MySQL: close db fail ...")
	}

	return nil
}
