package internal

import (
	"fmt"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/config"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/handler"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/migrations"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/repository"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *http.Server
}

func initLogger() {
	logger := log.New()
	logger.Out = os.Stdout
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func Inject() *gin.Engine {
	cfg := config.NewDbConfig()

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	log.Printf(connection)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Start migrating database \n")
	err = migrations.MigrationUp(db)
	if err != nil {
		log.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	router.Group("/").GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive"})
		return
	})

	return router
}

func (s *Server) Run(port string) error {
	initLogger()
	config.EnvsCheck()

	router := Inject()
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	log.Println("Shutting down")

	s.server.Close()
}
