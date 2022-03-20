package server

import (
	"context"
	"kienmatu/go-todos/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Configuration
	db     *gorm.DB
	logger *logrus.Logger
}

func NewServer(cfg *config.Configuration, db *gorm.DB, logger *logrus.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":" + s.cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Port)
	if err := s.echo.StartServer(server); err != nil {
		s.logger.Fatalln("Error starting Server: ", err)
	}

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Fatalln("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}