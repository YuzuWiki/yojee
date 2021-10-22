package pixiv

import "sync"

// engine
type Engine struct {
	lock sync.Mutex

	spiders map[string]interface{}
}

func (e *Engine) beforeStart() {}

func (e *Engine) Run() {}

func (e *Engine) afterClose() {}

func (e *Engine) Close() {}

func (e *Engine) RegisterSpider() {}
