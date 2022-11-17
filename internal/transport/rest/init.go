package rest

import (
	"net/http"

	"wlbot/internal/service/mikrotik"
	"wlbot/internal/service/notificator"

	"go.uber.org/zap"
)

type Server struct {
	l *zap.SugaredLogger

	*http.Server
	mikrotik    *mikrotik.Service
	notificator *notificator.Service
}

func NewServer(port string, mikrotik *mikrotik.Service) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: nil,
		},
		mikrotik: mikrotik,
	}
}

func (s *Server) setRoutes() {
	r := http.NewServeMux()
	r.HandleFunc("/send", s.CorsHandler(s.SendHandler()))
	r.HandleFunc("/wl", s.CorsHandler(s.AddIPHandler()))

	s.Handler = r
}
