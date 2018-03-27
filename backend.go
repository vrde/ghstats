package ghstats

import (
	"fmt"
)

type Serializable interface {
	Headers() []string
	Values() []interface{}
}

type Backend struct{}

func (b *Backend) Store(s Serializable) {
	fmt.Println(s.Values()...)
}

func GetBackend() *Backend {
	return &Backend{}
}
