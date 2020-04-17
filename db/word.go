package db

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
)

/*Words is slice of string */
type Words []string

/*GetWordsFromFile gives Words */
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

func (jStore *jsonStore) GetWords(ctx context.Context, limit int, allowSwear bool) (words Words, err error) {
	allWords, _ := jStore.GetSafeWords(ctx)

	if allowSwear == true {
		swearWords, _ := jStore.GetSwearWords(ctx)
		allWords = append(allWords, swearWords...)
	}

	allWords = shuffleWords(allWords)
	for i := 0; i < limit; i++ {
		words = append(words, allWords[i])
	}
	return
}
