package server

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	s.httpServer = httpServer
	return s.httpServer.ListenAndServe()
}
