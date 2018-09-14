package ghstats

type User struct {
	Id        int
	Login     string
	Name      string
	Bio       string
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}

type UsersBuffer []User

type Members struct {
	OrgId    int
	OrgLogin string
	buffer   UsersBuffer
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
	v := make([]interface{}, l*len(m.buffer))

	for i, x := range m.buffer {
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

func (m *Members) NewBuffer() Iterable {
	m.buffer = UsersBuffer{}
	return &m.buffer
}

func (m *UsersBuffer) Items() []interface{} {
	items := make([]interface{}, len(*m))
	for i, v := range *m {
		items[i] = v
	}
	return items
}
