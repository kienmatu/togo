package main

import (
	"log"

	"dangquang9a/go-location/config"
	"dangquang9a/go-location/db"
	"dangquang9a/go-location/internal/server"

	"github.com/sirupsen/logrus"
)

// @title Todo Application
// @description Location history application
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {

	log.Println("Starting api server")
	// Initialize config
	cfg := config.NewConfig()
	db := db.GetPostgresInstance(cfg, true)
	s := server.NewServer(cfg, db, logrus.New(), nil)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

}
