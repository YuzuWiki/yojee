package global

import (
	"os"
	"strconv"

	"github.com/YuzuWiki/yojee/module/job_pool"
)

var JobPool job_pool.IPool

func initPool() {
	if JobPool == nil {
		rate, err := strconv.Atoi(os.Getenv("POOL_RATE"))
		if err != nil {
			panic(err.Error())
		}

		JobPool = job_pool.New(rate)
	}
}
