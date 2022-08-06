package global

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dB *gorm.DB

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
	if dB == nil {
		db, err := gorm.Open("mysql", url())
		if err != nil {
			panic(fmt.Sprintf("MySQL: init db fail ... url=%s", url()))
		}
		dB = db
	}
	return dB
}

func NewDB() *sql.DB {
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
