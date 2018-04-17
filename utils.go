package ghstats

import (
	"errors"
)

type FetchStorer interface {
	Fetcher
	Storer
}

func FetchAndUpdate(a *API, b *Backend, f FetchStorer) error {
	for err := range a.FetchAll(f) {
		if err != nil {
			return err
		}
		b.Insert(f)
		u.Reset()
	}
}
