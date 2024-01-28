package server

import (
	"context"
	"net/http"
)

type Server struct {
	Http_Server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.Http_Server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.Http_Server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Http_Server.Shutdown(ctx)
}
