package db

import "context"

/*Storer Interface */
type Storer interface {
	GetSafeWords(context.Context) (Words, error)
	GetSwearWords(context.Context) (Words, error)
	GetWords(context.Context, int, bool) (Words, error)
}
