package server

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func Create(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	log.Printf("Stopping on http://localhost%s", serv.server.Addr)
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
