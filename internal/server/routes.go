package server

import (
	"net/http"

	"I_Dev_Kit/cmd/web"
	"I_Dev_Kit/cmd/web/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(pages.MainPage()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)
	s.RegisterTimeRoutes(r)
	s.RegisterFeatureRoutes(r)
	return r
}

func (s *Server) RegisterTimeRoutes(r *chi.Mux) {
	r.Get("/api/current-time", s.currentTimeHandler)
	r.Get("/api/current-date", s.currentDateHandler)
}

func (s *Server) RegisterFeatureRoutes(r *chi.Mux) {
	r.Get("/api/quick-stats", s.v.GetStats)
}
