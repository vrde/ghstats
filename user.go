package ghstats

import (
	"fmt"
)

const membersUrl = "/orgs/%s/members"

type User struct {
	Id        int
	Login     string
	Name      string
	Bio       string
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}

type Members struct {
	OrgId    int
	OrgLogin string
	Members  []User
}

func (m *Members) Table() Table {
	return Table{"members", []Column{
		Column{"id", "INTEGER PRIMARY KEY"},
		Column{"org_id", "INTEGER"},
		Column{"login", "TEXT"},
		Column{"name", "TEXT"},
		Column{"bio", "TEXT"},
		Column{"html_url", "TEXT"},
		Column{"avatar_url", "TEXT"}}}
}

func (m *Members) Values() []interface{} {
	l := len(m.Table().Columns)
	v := make([]interface{}, l*len(m.Members))

	for i, x := range m.Members {
		o := i * l
		v[o+0] = x.Id
		v[o+1] = m.OrgId
		v[o+2] = x.Login
		v[o+3] = x.Name
		v[o+4] = x.Bio
		v[o+5] = x.HtmlUrl
		v[o+6] = x.AvatarUrl
	}
	return v
}

func (m *Members) Url() string {
	return fmt.Sprintf(membersUrl, m.OrgLogin)
}

func (m *Members) Reset() interface{} {
	m.Members = &[]User{}
	return m.Members
}
