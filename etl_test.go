package ghstats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlTemplate(t *testing.T) {
	assert := assert.New(t)
	u := UrlTemplate{Name: "issuesUrl", Template: "repos/{{.Up.Up.Item.Login}}/{{.Up.Item.Name}}/issues?state=all"}
	s := Stack{&Issues{}, &Stack{&Repo{Name: "ghstats"}, &Stack{&Org{Login: "vrde_inc"}, nil}}}
	assert.Equal(u.Execute(&s), "repos/vrde_inc/ghstats/issues?state=all")
}
