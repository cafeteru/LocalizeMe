package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/impl"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func CreateServer(port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	config := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(config.Handler)
	initControllers(r)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &Server{server}
}

func (serv *Server) Close() error {
	slog.Debugf("Stopping on http://localhost%s", serv.server.Addr)
	return nil
}

func (serv *Server) Start() {
	slog.Debugf("Server running on http://localhost%s", serv.server.Addr)
	slog.Fatal(serv.server.ListenAndServe())
}

func initControllers(r *chi.Mux) {
	userPort := impl.CreateUserPort()
	userPort.InitRoutes(r)
}
