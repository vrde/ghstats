package ghstats

import (
	"fmt"
	"log"
	"sync"
)

const orgUrl = "/orgs/%s"

type Org struct {
	Id        int
	Login     string
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}

func (o *Org) Headers() []string {
	return []string{"id", "login", "html_url", "avatar_url"}
}

func (o *Org) Values() []interface{} {
	v := make([]interface{}, len(o.Headers()))
	v[0] = o.Id
	v[1] = o.Login
	v[2] = o.HtmlUrl
	v[3] = o.AvatarUrl
	return v
}

func update(ch <-chan error, b *Backend, s Serializable) error {
	for err := range ch {
		if err != nil {
			return err
		}
		b.Store(s)
	}
	return nil
}

func UpdateAllFromOrg(c *Context, b *Backend, name string) error {
	wg := sync.WaitGroup{}

	org := &Org{}
	err := update(c.FetchAll(fmt.Sprintf(orgUrl, name), org), b, org)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		members := &Members{}
		members.OrgId = org.Id
		err = update(c.FetchAll(fmt.Sprintf(membersUrl, name), &members.Members), b, members)
		if err != nil {
			log.Print(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		repos := &Repos{}
		repos.OrgId = org.Id

		for err := range c.FetchAll(fmt.Sprintf(reposUrl, name), &repos.Repos) {
			if err != nil {
				log.Print(err)
			}
			b.Store(repos)

			for _, repo := range repos.Repos {
				wg.Add(1)
				go func(id int, name string) {
					defer wg.Done()
					issues := &Issues{}
					issues.OrgId = org.Id
					issues.RepoId = id
					fmt.Printf(issuesUrl+"\n", org.Login, name)
					err = update(c.FetchAll(fmt.Sprintf(issuesUrl, org.Login, name), &issues.Issues), b, issues)
					if err != nil {
						log.Print(err)
					}
					b.Store(issues)
				}(repo.Id, repo.Name)
			}

		}
	}()

	wg.Wait()

	return nil
}
