package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
        "os"
        "io/ioutil"
        "github.com/vrde/gitstats/utils"
)

const IssueUrl = "https://api.github.com/repos/bigchaindb/bigchaindb/issues?state=closed"

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

func getIssues(url string) (Issues, error){
	resp, err := http.Get(url)

	if err != nil {
            return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
            return nil, err
	}

        content, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

	var issues Issues
        err = json.Unmarshal(content, &issues)
        if err != nil {
            return nil, err
	}
        linkHeader := utils.ParseLinkHeader(resp.Header.Get("Link"))

        return issues, linkHeader, err
}

func main() {
        issues, linkHeader, err := getIssues(IssueUrl)
        if err != nil {
            os.Exit(0)
        }

        fmt.Printf("%+v\n", issues)
        fmt.Println(linkHeader)
}
