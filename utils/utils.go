package utils

import (
    "strings"
)


func ParseLinkHeader(linkHeader string) map[string]string {
    links := make(map[string]string)

    for _, line := range strings.Split(linkHeader, ",") {
        line := strings.TrimSpace(line)
        tokens := strings.Split(line, ";")
        rel := strings.Trim(strings.Split(tokens[1], "=")[1], `"`)
        link := strings.Trim(tokens[0], "<>")
        links[rel] = link
    }

    return links
}
