// Package ghstats provides functions to extract issues from GitHub and the
// related types.
package ghstats

import (
	"strconv"
	"time"
)

// An array of issues of type Issue
type Issues []Issue

// A GitHub Issue
type Issue struct {
	Number      int
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ClosedAt    time.Time    `json:"closed_at"`
}

// A GitHub pull request
type PullRequest struct {
	Url string
}

// Thanks https://stackoverflow.com/a/33357286/597097 for the following idea.

// An Interface providing basic functions that help serializing a type in a
// CSV.
type CSVAble interface {
	GetHeaders() []string
	ToSlice() []string
}

// Extract the headers of an array of issues.
func (i *Issues) GetHeaders() []string {
	return []string{"number", "pr_url", "created_at", "updated_at", "closed_at"}
}

// Serialize multiple issues to an array of strings.
func (i *Issues) ToSlice() [][]string {
	acc := make([][]string, len(*i))

	for j, issue := range *i {
		acc[j] = issue.ToSlice()
	}

	return acc
}

// Extract the headers of an issue.
func (i *Issue) GetHeaders() []string {
	return []string{"number", "pr_url", "created_at", "updated_at", "closed_at"}
}

// Serialize an issue to an array of strings.
func (i *Issue) ToSlice() []string {
	url := ""

	if i.PullRequest != nil {
		url = i.PullRequest.Url
	}

	return []string{strconv.Itoa(i.Number), url, i.CreatedAt.String(), i.UpdatedAt.String(), i.ClosedAt.String()}
}
