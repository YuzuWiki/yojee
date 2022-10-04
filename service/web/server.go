package web

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/YuzuWiki/yojee/global"
)

type Server struct {
	router *gin.Engine

	isRun      bool
	listenPort int

	mu sync.Mutex
}

func (svr *Server) setupPProf() *Server {
	pprof.Register(svr.router, "yojee/pprof")
	return svr
}

func (svr *Server) setupLogger(logger *zerolog.Logger) *Server {
	svr.router.Use(func(ctx *gin.Context) {
		ctx.Next()
		logger.Info().Msg(fmt.Sprintf("[%s] %s", ctx.Request.Method, ctx.Request.URL.Path))
	})
	return svr
}

func (svr *Server) IsRun() bool {
	svr.mu.Lock()
	defer svr.mu.Unlock()

	return svr.isRun
}

func (svr *Server) shutdown() error {
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

func (svr *Server) Run() error {
	if svr.IsRun() {
		global.Logger.Error().Msg("HTTP server run: fail, server is run")
		return nil
	}

	svr.mu.Lock()
	defer svr.mu.Unlock()

	if err := svr.router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return err
	}

	global.Logger.Info().Msg("HTTP server run: ...")
	svr.isRun = true

	svr.setupLogger(global.Logger)
	svr.RegisterRoutes()
	if err := svr.router.Run(fmt.Sprintf(":%d", svr.listenPort)); err != nil {
		global.Logger.Err(err).Msg("HTTP server run: fail")
		return err
	}

	global.Logger.Info().Msg("HTTP server run: success")
	return nil
}

func Start(port string) {
	listenPort, err := strconv.Atoi(port)
	if err != nil {
		panic(err.Error())
	}

	svr := &Server{
		router:     gin.New(),
		isRun:      false,
		listenPort: listenPort,
		mu:         sync.Mutex{},
	}

	if err = svr.Run(); err != nil {
		panic("web service error")
	}
	return
}
