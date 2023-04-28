package global

import (
	"os"
	"strconv"
	"sync"

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

var (
	JobPool IPool

	jobOnce sync.Once
)

func initPool() {
	jobOnce.Do(func() {
		rate, err := strconv.Atoi(os.Getenv("POOL_RATE"))
		if err != nil {
			panic(err.Error())
		}

		JobPool = New(rate)
	})
}
