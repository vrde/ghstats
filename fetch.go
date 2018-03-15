package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vrde/gitstats/utils"
	"log"
	"net/http"
	"os"
	"time"
)

const IssueUrl = "https://api.github.com/repos/%s/issues?state=closed"

var Logger = log.New(os.Stderr, "", log.LstdFlags)

type Issues []Issue

type Issue struct {
	Number      int
	PullRequest PullRequest `json:"pull_request,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ClosedAt    time.Time   `json:"closed_at"`
}

type PullRequest struct {
	Url string
}

func fetchIssues(url string, issues *Issues) (error, string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
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
		os.Exit(1)
	}

	return utils.NextLinkHeader(resp.Header.Get("Link"))
}

func FetchIssues(repository string) *Issues {
	var issues Issues
	url := fmt.Sprintf(IssueUrl, repository)

	for {
		var (
			last Issues
			err  error
		)
		Logger.Println("Fetching", url)
		err, url = fetchIssues(url, &last)
		if err != nil {
			return &issues
		}
		issues = append(issues, last...)
	}
}

func main() {
	issues := FetchIssues(os.Args[1])
	spew.Dump(issues)
}
