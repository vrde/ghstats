package ghstats

type Org struct {
	Login  string
	buffer OrgBuffer
}

type OrgBuffer struct {
	Id         int
	Login      string
	MembersUrl string `json:"members_url"`
	ReposUrl   string `json:"repos_url"`
	HtmlUrl    string `json:"html_url"`
	AvatarUrl  string `json:"avatar_url"`
}

func (o *Org) Table() Table {
	return Table{"orgs", []Column{
		Column{"id", "INTEGER PRIMARY KEY"},
		Column{"login", "TEXT"},
		Column{"html_url", "TEXT"},
		Column{"avatar_url", "TEXT"}}}
}

func (o *Org) Values() []interface{} {
	v := make([]interface{}, len(o.Table().Columns))
	v[0] = o.buffer.Id
	v[1] = o.buffer.Login
	v[2] = o.buffer.HtmlUrl
	v[3] = o.buffer.AvatarUrl
	return v
}

func (o *Org) NewBuffer() Iterable {
	o.buffer = OrgBuffer{}
	return &o.buffer
}

func (o *OrgBuffer) Items() []interface{} {
	v := make([]interface{}, 1)
	v[0] = o
	return v
}
