package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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
	dateStr := now.Format("Monday, January 2")
	w.Write([]byte(dateStr))
}
