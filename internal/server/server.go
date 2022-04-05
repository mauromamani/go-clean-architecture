package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/config"
)

type server struct {
	cfg *config.Config
	gin *gin.Engine
}

func NewServer(cfg *config.Config) *server {
	return &server{
		gin: gin.Default(),
		cfg: cfg,
	}
}

// Run: run server on port 3000
func (s *server) Run() error {
	s.mapHandlers()

	if err := s.gin.Run(s.cfg.Server.Port); err != nil {
		return err
	}

	return nil
}
