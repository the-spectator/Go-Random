package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// const (
// 	versionHeader = "Accept"
// )

/*InitRouter initalizes the router */
func InitRouter(dep Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	// Version 1 API management
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router.HandleFunc("/words", wordHandler(dep)).Methods(http.MethodGet)
	// .Headers(versionHeader, v1)
	return
}
