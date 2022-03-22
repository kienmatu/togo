package main

import (
	"log"

	"kienmatu/go-todos/config"
	"kienmatu/go-todos/db"
	"kienmatu/go-todos/internal/server"

	"github.com/sirupsen/logrus"
)

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
