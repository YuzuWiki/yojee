package global

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func url() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
}

func InitDB() *gorm.DB {
	if db == nil {
		_db, err := gorm.Open("mysql", url())
		if err != nil {
			panic(fmt.Sprintf("MySQL: init db fail ... url=%s", url()))
		}
		db = _db
	}
	return db
}

func DB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db == nil {
		return nil
	}

	if err := db.Close(); err != nil {
		panic("MySQL: close db fail ...")
	}

	return nil
}
