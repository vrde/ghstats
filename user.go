package ghstats

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
	OrgId   int
	Members []User
}

func (m *Members) Headers() []string {
	return []string{"org_id", "user_id", "login", "name", "bio", "html_url", "avatar_url"}
}

func (m *Members) Values() []interface{} {
	l := len(m.Headers())
	v := make([]interface{}, l*len(m.Members))

	for i, x := range m.Members {
		o := i * l
		v[o+0] = m.OrgId
		v[o+1] = x.Id
		v[o+2] = x.Login
		v[o+3] = x.Name
		v[o+4] = x.Bio
		v[o+5] = x.HtmlUrl
		v[o+6] = x.AvatarUrl
	}
	return v
}
