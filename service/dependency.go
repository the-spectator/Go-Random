package service

import "go_random/db"

/*Dependencies of the random word generator */
type Dependencies struct {
	Store db.Storer
}
