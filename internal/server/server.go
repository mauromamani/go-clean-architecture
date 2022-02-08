package server

import "github.com/gin-gonic/gin"

type server struct {
	gin *gin.Engine
}

func NewServer() *server {
	return &server{gin: gin.Default()}
}

// Run: run server on port 3000
func (s *server) Run() error {
	s.mapHandlers()

	if err := s.gin.Run(":3000"); err != nil {
		return err
	}

	return nil
}
