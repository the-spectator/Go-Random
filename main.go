package main

import (
	"fmt"
	"go_random/config"
	"go_random/service"
	"strconv"

	"github.com/urfave/negroni"

	logger "github.com/sirupsen/logrus"
)

func main() {

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	config.Load()

	safeWords, err := service.GetWordsFromFile("words.json")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting words from words json")
		panic(err)
	}

	swearWords, err := service.GetWordsFromFile("swear.json")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting words from swear json")
		panic(err)
	}

	//declare dependency
	deps := service.Dependencies{
		SafeWords:  safeWords,
		SwearWords: swearWords,
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This can be changed to the service port number via environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}
