package main

import "net/http"

type Server struct {
	mux *http.ServeMux

	Port string
}

func (s *Server) MountHandlers() (err error) {
	mux := http.NewServeMux()

	//mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./static"))))
	mux.Handle("/webhook", &SyncHandler{})


	s.mux = mux
	return nil
}

func (s *Server) Start() (err error) {
	err = s.MountHandlers()
	err = http.ListenAndServe(":" + s.Port, s.mux)

	return err
}
