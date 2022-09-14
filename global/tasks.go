package global

import (
	"os"
	"strconv"

	"github.com/YuzuWiki/yojee/module/tasks"
)

var Pool tasks.IPool

func initPool() {
	if Pool == nil {
		rate, err := strconv.Atoi(os.Getenv("POOL_RATE"))
		if err != nil {
			panic(err.Error())
		}

		Pool = tasks.New(rate)
	}
}
