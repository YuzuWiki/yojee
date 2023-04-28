package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"os"
	"sync"
	"time"

	mysql2 "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	dbOnce sync.Once
)

var DATABASE = func() (database string) {
	database = os.Getenv("MYSQL_DATABASE")
	if len(database) == 0 {
		database = "yojee"
	}
	return
}

func dns() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		DATABASE(),
	)
}

func setConnectionPool() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(20)                  // 空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(50)                  // 打开数据库连接的最大数量
		sqlDB.SetConnMaxLifetime(60 * time.Minute) // 连接可复用的最大时间

		if err := sqlDB.Ping(); err != nil {
			panic(err)
		}
	}
}

func IsDuplicateEntry(err error) bool {
	errCode, _ := err.(*mysql2.MySQLError)

	return errCode.Number == 1062
}

func DB() *gorm.DB {
	return db
}

func InitDB() {
	dbOnce.Do(func() {
		_db, err := gorm.Open(
			mysql.New(mysql.Config{
				DSN:                       dns(),
				SkipInitializeWithVersion: false,
			}),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Error), // 日志输出级别
			},
		)
		if err != nil {
			panic(fmt.Sprintf("MySQL: init db fail ... dns=%s", dns()))
		}
		db = _db

		setConnectionPool()
	})
}
