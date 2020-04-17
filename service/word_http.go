package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

func wordHandler(dep Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			logger.WithField("err", err.Error()).Error("Error parsing query string")
			return
		}

		l := req.Form.Get("limit")
		limit := 5
		if l != "" {
			var err error
			limit, err = strconv.Atoi(l)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error converting limit params")
			}
		}

		swear := req.Form.Get("swear")
		allowSwear, _ := strconv.ParseBool(swear)
		words, err := dep.Store.GetWords(req.Context(), limit, allowSwear)

		if err != nil {
			panic(err)
		}

		respBytes, err := json.Marshal(words)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshalling ping response")
			rw.WriteHeader(http.StatusInternalServerError)
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}
