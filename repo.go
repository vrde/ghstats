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

func (r *Repos) Table() Table {
	return Table{"repos", []Column{
		Column{"id", "INTEGER PRIMARY KEY"},
		Column{"org_id", "INTEGER"},
		Column{"name", "TEXT"},
		Column{"full_name", "TEXT"},
		Column{"description", "TEXT"},
		Column{"html_url", "TEXT"},
		Column{"forks_count", "INTEGER"},
		Column{"stargazers_count", "INTEGER"},
		Column{"watchers_count", "INTEGER"},
	}}
}

func (r *Repos) Values() []interface{} {
	l := len(r.Table().Columns)
	v := make([]interface{}, l*len(r.Repos))

	for i, x := range r.Repos {
		o := i * l
		v[o+0] = x.Id
		v[o+1] = r.OrgId
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
