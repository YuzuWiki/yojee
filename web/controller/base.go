package controller

import (
	"sync"

	"github.com/rs/zerolog"
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
