package main

import (
	"fmt"
	"go_random/config"
	"go_random/db"
	"go_random/service"
	"os"

	"github.com/urfave/negroni"

	logger "github.com/sirupsen/logrus"
)

func main() {

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	config.Load()

	store, err := db.Init()
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	//declare dependency
	deps := service.Dependencies{
		Store: store,
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	// port := config.AppPort() // This can be changed to the service port number via environment variable.
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	server.Run(addr)
}
