package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HandlerFunc func(ctx context.Context, decode func(interface{}) error) (interface{}, error)

type Server struct {
	config *Config
	logger *log.Logger
	mux    *http.Server
}

func New(config *Config, logger *log.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
		mux:    &http.Server{Addr: config.Server.Address, Handler: nil},
	}
}

func (s *Server) Register() {
	mux := http.NewServeMux()
	for route, handler := range s.cnRoutes() {
		mux.HandleFunc(route, translateHandler(handler))
	}
	s.mux.Handler = mux
}

func (s *Server) Start() {
	s.logger.Println(fmt.Sprintf("Starting server on %s...", s.mux.Addr))

	go func() {
		if err := s.mux.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				s.logger.Println(fmt.Sprintf("Error during ListenAndServer: %v", err))
			}
		}
	}()

	// TODO: Should this just be its own binary?
	if s.config.StartOptions.Index {
		go func() { s.index() }()
	}
}

func (s *Server) Stop() {
	timeout := time.Duration(s.config.Server.Timeouts.Shutdown) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	s.logger.Println("Server shutting down...")
	if err := s.mux.Shutdown(ctx); err != nil {
		s.logger.Println(err)
	} else {
		s.logger.Println("Server successfully shut down!")
	}
}
