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
	db := db.GetPostgresInstance(cfg, false)
	s := server.NewServer(cfg, db, logrus.New())
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// log.Println("Application is running at ", conf.Port)

	// e.Logger.Fatal(e.Start(":"+conf.Port))
}
