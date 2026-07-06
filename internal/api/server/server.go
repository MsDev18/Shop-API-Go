package server

import (
	"fmt"
	"log"
	"shop/internal/api/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	R      *gin.Engine
	config Config
}

type Config struct {
	Host string `koanf:"host"`
	Port uint   `koanf:"port"`
}

func New(config Config) Server {
	r := gin.Default()

	router := router.New(r)
	router.Register()

	return Server{
		config: config,
		R:      r,
	}
}

func (s Server) Run() {
	log.Printf("start running server on %s:%d", s.config.Host, s.config.Port)
	err := s.R.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
