package global

import (
	"os"
	"strconv"

	"go.uber.org/ratelimit"
)

type TaskFunc func()

type IPool interface {
	Submit(TaskFunc)
}

func New(rate int) IPool {
	p := pool{limiter: ratelimit.New(rate)}
	return &p
}

type pool struct {
	limiter ratelimit.Limiter
}

func (p *pool) Submit(fn TaskFunc) {
	p.limiter.Take()
	fn()
}

var JobPool IPool

func initPool() {
	if JobPool == nil {
		rate, err := strconv.Atoi(os.Getenv("POOL_RATE"))
		if err != nil {
			panic(err.Error())
		}

		JobPool = New(rate)
	}
}
