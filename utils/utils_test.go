package utils

import (
	"testing"
)

func TestLinkHeaderParses(t *testing.T) {
	links := ParseLinkHeader(
		`<https://api.github.com/resource?page=2>; rel="next",
         <https://api.github.com/resource?page=5>; rel="last"`)

	expected := map[string]string{
		"next": "https://api.github.com/resource?page=2",
		"last": "https://api.github.com/resource?page=5",
	}

	for expectedKey, expectedValue := range expected {
		if links[expectedKey] != expectedValue {
			t.Error(expectedKey, expectedValue, links[expectedKey])
		}
	}
}
