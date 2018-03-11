package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vrde/gitstats/utils"
	"net/http"
	"os"
	"time"
)

const IssueUrl = "https://api.github.com/repos/bigchaindb/bigchaindb/issues?state=closed"

type Issues []*Issue

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

func main() {
	resp, err := http.Get(IssueUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "GET failed: %v\n", err)
		os.Exit(1)
	}

	var result Issues

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", IssueUrl, err)
		os.Exit(1)
	}

	spew.Dump(utils.ParseLinkHeader(resp.Header.Get("Link")))

	// fmt.Printf("%+v\n", result)
}
