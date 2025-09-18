package httpsh

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"grep/2.18/domain"
)

type Server struct {
	port     string
	logger   domain.Logger
	service  domain.EventService
	storage  domain.Storage
	handlers domain.EventHandler
	srv      *http.Server
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewServer(port string, logger domain.Logger, service domain.EventService, storage domain.Storage, handlers domain.EventHandler) *Server {
	return &Server{
		port:     port,
		logger:   logger,
		service:  service,
		storage:  storage,
		handlers: handlers,
	}
}

func (s *Server) RunServ() error {
	if s.port == "" {
		s.logger.Error("Port is not set")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	RegisterRoutes(mux, s.handlers)
	mux.Handle("/", &handleDef{})

	srv := &http.Server{
		Addr:         "0.0.0.0:" + s.port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	s.srv = srv

	s.logger.Info("Starting server", "port", s.port)
	return s.srv.ListenAndServe()
}

type handleDef struct{}

func (h *handleDef) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusInternalServerError
	if r.Method == "OPTIONS" {
		statusCode = http.StatusOK
	}

	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Message: "Undefined Error, please check your method or endpoint correctness",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
