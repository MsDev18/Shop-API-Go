package server

import (
	"fmt"
	"log"
	"net/http"
	"shop/internal/api/router"
	"shop/internal/pkg/richerror"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     Config
	httpServer *http.Server
	// handler statements
	handlers []router.RegisterRoutes
}

type Config struct {
	Host         string `koanf:"host"`
	Port         uint   `koanf:"port"`
	WriteTimeout uint   `koanf:"write_timeout"`
	ReadTimeout  uint   `koanf:"read_timeout"`
	IdleTimeout  uint   `koanf:"idle_timeout"`
	Env          string `koanf:"env"`
}

func New(config Config, handlers ...router.RegisterRoutes) Server {
	// validation env
	env := Env(config.Env)
	if !env.IsValid() {
		log.Fatalf("Invalid environment: %s", config.Env)
	}
	// validate timeout values
	if config.ReadTimeout == 0 || config.WriteTimeout == 0 || config.IdleTimeout == 0 {
		log.Fatalf("Invalid timeout values: read_timeout=%d, write_timeout=%d, idle_timeout=%d", config.ReadTimeout, config.WriteTimeout, config.IdleTimeout)
	}
	
	// set from gin mode based on env
	if env == EnvDevelopment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// create gin engine
	engine := gin.Default()
	// register routes
	appRouter := router.New(engine, handlers...)
	appRouter.Register()

	// manually create http server to set timeouts
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:      engine,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
	}

	return Server{
		config:     config,
		httpServer: httpServer,
		// handlers statements
		handlers: handlers,
	}
}

func (s Server) Run() error {
	const op = "server.Run"
	log.Printf("start running server on %s:%d", s.config.Host, s.config.Port)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return richerror.New().
			SetOp(op).
			SetMsg("failed to run server").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	return nil
}
