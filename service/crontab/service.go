package crontab

import (
	"github.com/YuzuWiki/yojee/global"
)

func registerTask() map[string]func() {
	return map[string]func(){}
}

func Start() {
	for spec, fn := range registerTask() {
		if _, err := global.Crontab.AddFunc(spec, fn); err != nil {
			panic(err.Error())
		}
	}

	global.Crontab.Start()
}
