package server

import (
	"fmt"
	"github.com/startup-of-zero-reais/takeout-users-api/server/infra/adapter"
)

type Server struct {
	fw *adapter.Adapter
}

func NewServer() *Server {
	fw := adapter.NewServer()

	return &Server{
		fw: fw,
	}
}

func (s *Server) Listen(port ...string) {
	s.Router()

	if len(port) != 1 {
		s.fw.Logger.Fatal(s.fw.Start(":8080"))
	} else {
		runPort := fmt.Sprintf(":%s", port[0])
		s.fw.Logger.Fatal(s.fw.Start(runPort))
	}
}
