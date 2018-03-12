package utils

import (
	"strings"
)

type LinkHeader struct {
    Url string
    Rel string
}

type LinkHeaders []LinkHeader

func ParseLinkHeader(linkHeader string) LinkHeaders {
	var links LinkHeaders

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

                linkHeader := LinkHeader{link, rel}
                links = append(links, linkHeader)
	}

	return links
}
