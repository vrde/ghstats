package ghstats

import (
	"testing"
)

func TestParseLinkHeader(t *testing.T) {
	tests := []struct {
		header string
		next   string
	}{
		{
			`<https://api.github.com/resource?page=2>; rel="next",\n<https://api.github.com/resource?page=5>; rel="last"`,
			"https://api.github.com/resource?page=2",
		},
		{
			`<https://api.github.com/resource?page=2>; rel="next",\n<https://api.github.com/resource?page=5> rel="last"`,
			"https://api.github.com/resource?page=2",
		},
		{
			`<https://api.github.com/resource?page=5> rel="last"`,
			"",
		},
	}

	for _, test := range tests {
		next := nextLinkHeader(test.header)

		if next != test.next {
			t.Errorf("nextLinkHeader(%q)\nGot:  %v\nWant: %v", test.header, next, test.next)
		}
	}
}
