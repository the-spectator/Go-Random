package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

type jsonStore struct {
	SafeWords  Words
	SwearWords Words
}

func (jStore *jsonStore) GetSafeWords(context.Context) (Words, error) {
	return jStore.SafeWords, nil
}

func (jStore *jsonStore) GetSwearWords(context.Context) (Words, error) {
	return jStore.SwearWords, nil
}

/*Init Json Store*/
func Init() (s Storer, err error) {
	safeWords, err := getWordsFromFile("words.json")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting words from words json")
		panic(err)
	}

	swearWords, err := getWordsFromFile("swear.json")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting words from swear json")
		panic(err)
	}

	jStore := jsonStore{
		SafeWords:  safeWords,
		SwearWords: swearWords,
	}

	s = &jStore
	return
}
