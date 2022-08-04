package global

import (
	_cron "github.com/robfig/cron"
)

var Cron *_cron.Cron

func InitCron() {
	if Cron == nil {
		Cron = _cron.New()
	}
}
