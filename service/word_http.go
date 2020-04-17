package service

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	logger "github.com/sirupsen/logrus"
)

type Words []string

func getWordsFromFile(name string) (words Words, err error) {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error opening File")
		return
	}

	err = json.NewDecoder(file).Decode(&words)
	return
}

func shuffleWords(words Words) Words {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	return words
}

func getWords(limit int, allowSwear bool) (words Words, err error) {
	allWords := Words{}

	if allowSwear == true {
		var swearWords Words
		swearWords, err = getWordsFromFile("swear.json")

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error getting words from swear json")
			return
		}

		allWords = append(allWords, swearWords...)
	}

	allWords, err = getWordsFromFile("words.json")

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting words from words json")
		return
	}

	allWords = shuffleWords(allWords)
	for i := 0; i < limit; i++ {
		words = append(words, allWords[i])
	}
	return
}

func wordHandler(rw http.ResponseWriter, req *http.Request) {
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
	words, err := getWords(limit, allowSwear)

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
}
