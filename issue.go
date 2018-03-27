// Package ghstats provides functions to extract issues from GitHub and the
// related types.
package ghstats

import (
	//"strconv"
	"time"
)

const issuesUrl = "/repos/%s/%s/issues"

type Issues struct {
	OrgId  int
	RepoId int
	Issues []Issue
}

// A GitHub Issue
type Issue struct {
	Id          int
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

func (i *Issues) Headers() []string {
	return []string{"org_id", "repo_id", "issue_id", "number", "pr_url", "created_at", "updated_at", "closed_at"}
}

func (i *Issues) Values() []interface{} {
	l := len(i.Headers())
	v := make([]interface{}, l*len(i.Issues))

	for j, x := range i.Issues {
		o := j * l
		v[o+0] = i.OrgId
		v[o+1] = i.RepoId
		v[o+2] = x.Id
		v[o+3] = x.Number
		v[o+4] = x.PullRequest.Url
		v[o+5] = x.CreatedAt
		v[o+6] = x.UpdatedAt
		v[o+7] = x.ClosedAt
	}
	return v
}
