package controller

import (
	"github.com/rs/zerolog"
	"sync"
)

type baseController struct {
	_log *zerolog.Logger
	once sync.Once
}

func (ctrl *baseController) log() *zerolog.Logger {
	ctrl.once.Do(func() {
		if ctrl._log == nil {
			*ctrl._log = zerolog.Logger{}
		}
	})
	return ctrl._log
}
