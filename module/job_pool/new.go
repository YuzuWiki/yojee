package job_pool

import (
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
