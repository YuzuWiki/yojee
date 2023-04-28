package global

import (
	_cron "github.com/robfig/cron/v3"
	"sync"
)

var (
	Crontab *_cron.Cron

	crontabOnce sync.Once
)

func InitCron() {
	crontabOnce.Do(func() {
		Crontab = _cron.New()
	})
}
