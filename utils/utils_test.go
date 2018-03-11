package utils

import (
	"testing"
)

func compare(t *testing.T, returned map[string]string, expected map[string]string) {
	for expectedKey, expectedValue := range expected {
		if returned[expectedKey] != expectedValue {
			t.Error(expectedKey, expectedValue, returned[expectedKey])
		}
	}
}

func TestParseLinkHeaderParsesGoodHeaders(t *testing.T) {
	links := ParseLinkHeader(
		`<https://api.github.com/resource?page=2>; rel="next",
         <https://api.github.com/resource?page=5>; rel="last"`)

	expected := map[string]string{
		"next": "https://api.github.com/resource?page=2",
		"last": "https://api.github.com/resource?page=5",
	}

	compare(t, links, expected)
}

func TestParseLinkHeaderReturnsAnEmptyMapOnEmptyStrig(t *testing.T) {
	links := ParseLinkHeader("")
	expected := map[string]string{}
	compare(t, links, expected)
}

func TestLinkHeaderParsesInvalidHeaders(t *testing.T) {
	links := ParseLinkHeader(
		`<https://api.github.com/resource?page=2>; rel="next",
         <https://api.github.com/resource?page=5> "last"`)

	expected := map[string]string{
		"next": "https://api.github.com/resource?page=2",
	}

	compare(t, links, expected)
}
