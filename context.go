package ghstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vrde/ghstats/utils"
	"log"
	"net/http"
	"os"
)

// The context to use when doing HTTP requests.
//
// It contains the GitHub authentication token and the API root.
type Context struct {
	GitHubToken   string
	GitHubRootAPI string
}

func GetContext() *Context {
	ctx := Context{}
	ctx.GitHubRootAPI = "https://api.github.com"

	if token, defined := os.LookupEnv("GITHUB_TOKEN"); defined {
		ctx.GitHubToken = token
	} else {
		log.Fatal("GITHUB_TOKEN env variable not found.")
	}

	return &ctx
}

func (c *Context) fetch(url string, v interface{}) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+c.GitHubToken)

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("error fetching %s", url)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("error retrieving <%s>: %v", url, resp.Status))
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return "", err
	}

	log.Printf("got %s", url)

	return utils.NextLinkHeader(resp.Header.Get("Link")), nil
}

func (c *Context) Fetch(url string, v interface{}) (string, error) {
	url = c.GitHubRootAPI + url
	return c.fetch(url, v)
}

func (c *Context) FetchAll(url string, v interface{}) <-chan error {
	url = c.GitHubRootAPI + url
	ch := make(chan error)
	go func() {
		defer close(ch)
		for url != "" {
			next, err := c.fetch(url, v)
			url = next

			if err != nil {
				ch <- err
			} else {
				ch <- nil
			}
		}
	}()
	return ch
}
