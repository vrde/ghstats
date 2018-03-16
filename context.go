package ghstats

// The context to use when doing HTTP requests.
//
// It contains the GitHub token to authenticate and query the API.
type Context struct {
	GitHubToken string
}
