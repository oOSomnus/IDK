package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

	return r
}

func (s *Server) RegisterTimeRoutes(r *chi.Mux) {
	r.Get("/api/current-time", s.currentTimeHandler)
	r.Get("/api/current-date", s.currentDateHandler)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	resp := time.Now().Format("15:04:05")
	w.Write([]byte(resp))
}

func (s *Server) currentDateHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	// 格式化为 "Monday, January 1" 形式
	dateStr := now.Format("Monday, January 2")
	w.Write([]byte(dateStr))
}
