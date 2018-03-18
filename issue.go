// Package ghstats provides functions to extract issues from GitHub and the
// related types.
package ghstats

import (
	"strconv"
	"time"
)

// Interface to serialize a struct to an array of strings.
type Slicer interface {
	ToSlice() []string
}

// Headers for an issue. Used for serialization in conjunction with the
// interface Slicer.
var IssueHeaders = []string{"number", "pr_url", "created_at", "updated_at", "closed_at"}

// An array of issues of type Issue
type Issues []Issue

// A GitHub Issue
type Issue struct {
	Number      int
	PullRequest PullRequest `json:"pull_request,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ClosedAt    time.Time   `json:"closed_at"`
}

// A GitHub pull request
type PullRequest struct {
	Url string
}

// Serialize multiple issues to an array of strings.
func (i *Issues) ToSlice() [][]string {
	acc := make([][]string, len(*i))

	for j, issue := range *i {
		acc[j] = issue.ToSlice()
	}

	return acc
}

// Serialize an issue to an array of strings.
func (i *Issue) ToSlice() []string {
	url := ""

	if i.PullRequest == (PullRequest{}) {
		url = i.PullRequest.Url
	}

	return []string{strconv.Itoa(i.Number), url, i.CreatedAt.String(), i.UpdatedAt.String(), i.ClosedAt.String()}
}
