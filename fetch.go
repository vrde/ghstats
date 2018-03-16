package ghstats

import (
	"encoding/json"
	"fmt"
	"github.com/vrde/ghstats/utils"
	"log"
	"net/http"
	"os"
)

const IssueUrl = "https://api.github.com/repos/%s/issues?state=closed"

// Fetch issues from a specific URL.
func fetchIssues(ctx *Context, url string, issues *Issues) (error, string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+ctx.GitHubToken)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "GET failed with status: %d (%v)\n", resp.StatusCode, err)
		os.Exit(1)
	}

	if err := json.NewDecoder(resp.Body).Decode(issues); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		fmt.Println(resp.Body)
		os.Exit(1)
	}

	return utils.NextLinkHeader(resp.Header.Get("Link"))
}

// Fetch all the issues from a GitHub repository, and send them to a channel.
func FetchIssues(ctx *Context, repository string, ch chan<- *Issues) {
	url := fmt.Sprintf(IssueUrl, repository)

	for {
		var (
			issues Issues
			err    error
		)
		log.Println("Fetching", url)
		err, url = fetchIssues(ctx, url, &issues)
		if err != nil {
			close(ch)
			return
		}
		ch <- &issues
	}
}
