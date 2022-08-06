package global

import (
	_cron "github.com/robfig/cron"
)

var Crontab *_cron.Cron

func InitCron() {
	if Crontab == nil {
		Crontab = _cron.New()
	}
}
