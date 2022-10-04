package global

import (
	_cron "github.com/robfig/cron/v3"
)

var Crontab *_cron.Cron

func InitCron() {
	if Crontab == nil {
		Crontab = _cron.New()
	}
}
