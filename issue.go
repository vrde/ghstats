package gitstats

import (
	"strconv"
	"time"
)

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

// Thanks https://stackoverflow.com/a/33357286/597097
type CSVAble interface {
	GetHeaders() []string
	ToSlice() []string
}

func (i *Issues) GetHeaders() []string {
	return []string{"number", "pr_url", "created_at", "updated_at", "closed_at"}
}

func (i *Issues) ToSlice() [][]string {
	acc := make([][]string, len(*i))
	for j, issue := range *i {
		acc[j] = issue.ToSlice()
	}
	return acc
}

func (i *Issue) GetHeaders() []string {
	return []string{"number", "pr_url", "created_at", "updated_at", "closed_at"}
}

func (i *Issue) ToSlice() []string {
	url := ""
	if i.PullRequest != nil {
		url = i.PullRequest.Url
	}
	return []string{strconv.Itoa(i.Number), url, i.CreatedAt.String(), i.UpdatedAt.String(), i.ClosedAt.String()}
}
