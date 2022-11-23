package rest

import (
	"net/http"
	"time"

	"wlbot/internal/service/mikrotik"
	"wlbot/internal/service/notificator"

	"go.uber.org/zap"
)

type Server struct {
	l *zap.SugaredLogger

	*http.Server
	mikrotik    mikrotik.Provider
	notificator *notificator.Service
}

func NewServer(l *zap.SugaredLogger, port string, n *notificator.Service, m mikrotik.Provider) *Server {
	const readTimeout = time.Second * 5

	httpSrv := &http.Server{
		Addr:              ":" + port,
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: readTimeout,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	s := &Server{
		l:           l,
		Server:      httpSrv,
		mikrotik:    m,
		notificator: n,
	}

	s.setRoutes()

	return s
}

func (s *Server) setRoutes() {
	r := http.NewServeMux()
	r.HandleFunc("/send", s.Cors(s.SendHandler()))
	r.HandleFunc("/wl", s.Cors(s.AddIPHandler()))

	s.Handler = r
}
