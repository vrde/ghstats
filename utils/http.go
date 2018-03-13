package utils

import (
	"errors"
	"strings"
)

func NextLinkHeader(linkHeader string) (error, string) {
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
			return nil, link
		}
	}

	return errors.New(`cannot find "next" link in header`), ""
}
