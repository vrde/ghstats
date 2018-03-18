package utils

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
		next := NextLinkHeader(test.header)

		if next != test.next {
			t.Errorf("NextLinkHeader(%q)\nGot:  %v\nWant: %v", test.header, next, test.next)
		}
	}
}
