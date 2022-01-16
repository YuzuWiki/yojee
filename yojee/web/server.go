package web

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/gin-contrib/pprof"
	"github.com/like9th/yojee/yojee/global"
)


type Server struct {
	*gin.Engine

	isRun bool
	listenPort int

	mu sync.Mutex
}

func (svr *Server) setupPProf() *Server {
	pprof.Register(svr.Engine, "yojee/pprof")
	return svr
}

func (svr *Server) setupLogger(logger *zerolog.Logger) *Server {
	svr.Use(func(ctx *gin.Context) {
			ctx.Next()
			logger.Info().Msg(fmt.Sprintf("[%s] %s", ctx.Request.Method, ctx.Request.URL.Path))
	})
	return svr
}

func (svr *Server) IsRun() bool  {
	svr.mu.Lock()
	defer svr.mu.Unlock()

	return svr.isRun
}

func (svr *Server) shutdown() error  {
	if !svr.IsRun() {
		return nil
	}

	svr.mu.Lock()
	defer svr.mu.Unlock()

	global.Logger.Info().Msg("HTTP server shutdown: ....")
	svr.isRun = false
	// do something
	global.Logger.Info().Msg("HTTP server shutdown: success")
	return nil
}

func (svr *Server) Shutdown() error {
	return svr.shutdown()
}

func (svr *Server) Run() error  {
	if svr.IsRun(){
		global.Logger.Error().Msg("HTTP server run: fail, server is run")
		return nil
	}

	svr.mu.Lock()
	defer svr.mu.Unlock()

	global.Logger.Info().Msg("HTTP server run: ...")
	svr.isRun = true

	svr.setupLogger(global.Logger)
	svr.RegisterRoutes()
	if err := svr.Engine.Run(fmt.Sprintf(":%d", svr.listenPort)); err != nil {
		global.Logger.Err(err).Msg("HTTP server run: fail")
		return err
	}

	global.Logger.Info().Msg("HTTP server run: success")
	return nil
}


func NewServer(listenPort int) *Server {
	global.InitLogger()
	svr := &Server{
		Engine:     gin.New(),
		isRun:      false,
		listenPort: listenPort,
		mu:         sync.Mutex{},
	}
	return svr
}
