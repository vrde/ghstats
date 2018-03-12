package main

import (
	"encoding/json"
	"fmt"
	"github.com/vrde/gitstats/utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const IssueUrl = "https://api.github.com/repos/bigchaindb/bigchaindb/issues?state=closed"
const auth = ""

type Issues []Issue

type Issue struct {
	Number      int
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ClosedAt    time.Time    `json:"closed_at"`
}

type PullRequest struct {
	Url string
}

type LinkHeader struct {
	Url string
	Rel string
}

type LinkHeaders []LinkHeader

func getIssues(url string, auth string) (Issues, map[string]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Authorization", "token "+auth)
	resp, err := client.Do(req)

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var issues Issues
	err = json.Unmarshal(content, &issues)
	if err != nil {
		return nil, nil, err
	}
	linkHeader := utils.ParseLinkHeader(resp.Header.Get("Link"))

	return issues, linkHeader, err
}

func main() {
	var issues Issues
	var issuesPage Issues
	var linkHeaders map[string]string
	var err error
	url := IssueUrl

	for url != linkHeaders["last"] {
		// Make this a go routine
		issuesPage, linkHeaders, err = getIssues(url, auth)
		url = linkHeaders["next"]
		issues = append(issues, issuesPage...)
		fmt.Println("Downloading: " + url)
	}
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("%+v\n", len(issues))
}
