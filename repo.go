package ghstats

const reposUrl = "/orgs/%s/repos"

type Repos struct {
	OrgId int
	Repos []Repo
}

type Repo struct {
	Id              int
	Name            string
	FullName        string `json:"full_name"`
	Description     string
	HtmlUrl         string `json:"html_url"`
	ForksCount      int    `json:"forks_count"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
}

func (r *Repos) Headers() []string {
	return []string{"org_id", "repo_id", "name", "full_name", "description", "html_url", "forks_count", "stargazers_count", "whatchers_count"}
}

func (r *Repos) Values() []interface{} {
	l := len(r.Headers())
	v := make([]interface{}, l*len(r.Repos))

	for i, x := range r.Repos {
		o := i * l
		v[o+0] = r.OrgId
		v[o+1] = x.Id
		v[o+2] = x.Name
		v[o+3] = x.FullName
		v[o+4] = x.Description
		v[o+5] = x.HtmlUrl
		v[o+6] = x.ForksCount
		v[o+7] = x.StargazersCount
		v[o+8] = x.WatchersCount
	}
	return v
}
