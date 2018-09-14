package ghstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Fetcher interface {
	Url() string
	Reset() interface{}
}

// The context to use when doing HTTP requests.
//
// It contains the GitHub authentication token and the API root.
type API struct {
	GitHubToken   string
	GitHubRootAPI string
}

func GetAPI() *API {
	api := API{}
	api.GitHubRootAPI = "https://api.github.com"

	if token, defined := os.LookupEnv("GITHUB_TOKEN"); defined {
		api.GitHubToken = token
	} else {
		log.Fatal("GITHUB_TOKEN env variable not found.")
	}

	return &api
}

func (c *API) Fetch(url string, v interface{}) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+c.GitHubToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("error fetching %s", url)
		return "", err
	}

	defer resp.Body.Close()

	log.Printf("GET [%d]: %s, %p", resp.StatusCode, url, v)
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("error retrieving <%s>: %v", url, resp.Status))
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return "", err
	}

	return nextLinkHeader(resp.Header.Get("Link")), nil
}

// Extract the "next" link from the Headers of an HTTP request.
// If no "next" link is found, return the empty string.
func nextLinkHeader(linkHeader string) string {
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
