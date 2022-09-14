package tasks

import (
	"go.uber.org/ratelimit"
)

type pool struct {
	limiter ratelimit.Limiter
}

func (p *pool) Submit(fn TaskFunc) {
	p.limiter.Take()
	fn()
}
