package utils

import (
	"strings"
)

// Extract the "next" link from the Headers of an HTTP request.
// If no "next" link is found, return the empty string.
func NextLinkHeader(linkHeader string) string {
	for _, line := range strings.Split(linkHeader, ",") {
		line := strings.TrimSpace(line)

		linkTokens := strings.Split(line, ";")
		if len(linkTokens) != 2 {
			continue
		}
		link := strings.Trim(linkTokens[0], "<>")

		relTokens := strings.Split(linkTokens[1], "=")
		if len(relTokens) != 2 {
			continue
		}
		rel := strings.Trim(relTokens[1], `"`)

		if rel == "next" {
			return link
		}
	}
	return ""
}
