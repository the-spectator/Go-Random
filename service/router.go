package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// const (
// 	versionHeader = "Accept"
// )

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()

	// Version 1 API management
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router.HandleFunc("/words", wordHandler).Methods(http.MethodGet)
	// .Headers(versionHeader, v1)
	return
}
