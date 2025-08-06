package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	v1 "I_Dev_Kit/internal/controller/v1"
	"I_Dev_Kit/internal/database"
	"I_Dev_Kit/internal/repo"
	"I_Dev_Kit/internal/usecase/feature"
	"I_Dev_Kit/internal/usecase/project"
	"I_Dev_Kit/pkg/logger"
)

type Server struct {
	port int

	db database.Service

	v *v1.V1
}

func initControllers(db database.Service) *v1.V1 {
	l := logger.New("debug")
	featureRepo := repo.NewFeatureRepo(db.DB())
	projectRepo := repo.NewProjectRepo(db.DB())
	p := project.New(projectRepo)
	f := feature.New(featureRepo)
	v := v1.New(p, f, l)
	return v
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	v := initControllers(db)
	NewServer := &Server{
		port: port,
		db:   db,
		v:    v,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
