package utils

import (
	"fmt"
	"testing"
)

func TestParseLinkHeader(t *testing.T) {
	tests := []struct {
		header string
		next   string
		err    bool
	}{
		{
			`<https://api.github.com/resource?page=2>; rel="next",\n<https://api.github.com/resource?page=5>; rel="last"`,
			"https://api.github.com/resource?page=2",
			false,
		},
		{
			`<https://api.github.com/resource?page=2>; rel="next",\n<https://api.github.com/resource?page=5> rel="last"`,
			"https://api.github.com/resource?page=2",
			false,
		},
		{
			`<https://api.github.com/resource?page=5> rel="last"`,
			"",
			true,
		},
	}

	for _, test := range tests {
		next, err := NextLinkHeader(test.header)
		call := fmt.Sprintf("NextLinkHeader(%q)", test.header)

		if test.err && err == nil || !test.err && err != nil {
			t.Errorf("%s\nGot error: %v\nExpecting error: %t", call, err, test.err)
		} else if next != test.next {
			t.Errorf("NextLinkHeader(%q)\nGot:  %v\nWant: %v", test.header, next, test.next)
		}
	}
}
