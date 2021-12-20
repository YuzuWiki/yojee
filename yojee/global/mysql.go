package global

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	once.Do(func() {
		if DB == nil {
			db, err := 	gorm.Open("")
			if err != nil {
				panic("MySQL: init db fail ...")
			}

			DB = db
		}
	})

	return DB
}

func CloseDB() error  {
	if DB == nil {
		return nil
	}

	once.Do(func() {
		if err := DB.Close(); err != nil {
			panic("MySQL: close db fail ...")
		}
	})

	return nil
}


