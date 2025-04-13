package rest

import (
	"context"
	"net/http"
	"time"
)

type serverREST struct {
	httpServer *http.Server
}

// NewServerREST конструктор для serverREST.
func NewServerREST() *serverREST {
	return &serverREST{
		httpServer: &http.Server{
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

// Run запускает REST server.
func (s *serverREST) Run(port string, handler http.Handler) error {
	s.httpServer.Addr = ":" + port
	s.httpServer.Handler = handler

	return s.httpServer.ListenAndServe()
}

// ShutDown останавливает REST сервер.
func (s *serverREST) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
