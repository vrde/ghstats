package ghstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vrde/ghstats/utils"
	"log"
	"net/http"
)

const IssueUrl = "https://api.github.com/repos/%s/issues?state=closed"

type IssuesResponse struct {
	Issues *Issues
	Url    string
	Error  error
}

// Fetch issues from a URL and return the next URL to follow for even moar
// issues.
func fetchIssues(ctx *Context, url string, issues *Issues) (error, string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+ctx.GitHubToken)
	resp, err := client.Do(req)

	if err != nil {
		return err, ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status), ""
	}

	if err := json.NewDecoder(resp.Body).Decode(issues); err != nil {
		return err, ""
	}

	return nil, utils.NextLinkHeader(resp.Header.Get("Link"))
}

// Fetch all the issues from a GitHub repository, and send them to a channel.
func FetchIssues(ctx *Context, repository string, ch chan<- *IssuesResponse) {
	var err error
	url := fmt.Sprintf(IssueUrl, repository)

	for {
		i := &IssuesResponse{}
		i.Url = url

		log.Println("Fetching", url)
		err, url = fetchIssues(ctx, url, i.Issues)
		i.Error = err

		ch <- i

		if err != nil {
			close(ch)
			return
		}

		if url == "" {
			close(ch)
			return
		}

	}
}
