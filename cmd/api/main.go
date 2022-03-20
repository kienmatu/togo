package main

import (
	"kienmatu/go-todos/config"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Initialize config
	conf := config.NewConfig()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	log.Println("Application is running at ", conf.Port)

	e.Logger.Fatal(e.Start(":"+conf.Port))

}