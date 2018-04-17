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

func (o *Org) Table() Table {
	return Table{"orgs", []Column{
		Column{"id", "INTEGER PRIMARY KEY"},
		Column{"login", "TEXT"},
		Column{"html_url", "TEXT"},
		Column{"avatar_url", "TEXT"}}}
}

func (o *Org) Values() []interface{} {
	v := make([]interface{}, len(o.Table().Columns))
	v[0] = o.Id
	v[1] = o.Login
	v[2] = o.HtmlUrl
	v[3] = o.AvatarUrl
	return v
}

func (o *Org) Url() string {
	return fmt.Sprintf(orgUrl, o.Login)
}

func (o *Org) Reset() interface{} {
	return o
}

func update(ch <-chan error, b *Backend, s Storer) error {
	for err := range ch {
		if err != nil {
			return err
		}
		if err = b.Insert(s); err != nil {
			return err
		}
	}
	return nil
}

func UpdateAllFromOrg(a *API, b *Backend, name string) error {
	wg := sync.WaitGroup{}

	org := &Org{Login: name}
	err := FetchAndUpdate(a, b, org)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		members := &Members{}
		members.OrgId = org.Id
		err = FetchAndUpdate(a, b, members)
		if err != nil {
			log.Printf("error updating org members: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r := &Repos{OrgId: org.Id, OrgLogin: org.Login}

		for err := range a.FetchAll(r) {
			if err != nil {
				log.Printf("error updating org members: %v", err)
			}
			b.Insert(r)
			for _, repo := range r.Repos {
				wg.Add(1)
				go func(orgId int, repoId int) {
					defer wg.Done()
					i := &Issues{OrgId: orgId, RepoId: repo.Id, RepoName: repo.RepoName}
					FetchAndUpdate(a, b, i)
				}(repo.Id, repo.Name)
			}
			u.Reset()
		}

		for err := range a.FetchAll(fmt.Sprintf(reposUrl, name), &repos.Repos) {
			if err != nil {
				log.Printf("error fetching repo <%s>: %v", name, err)
			}
			if err = b.Insert(repos); err != nil {
				log.Printf("error updating repo <%s>: %v", name, err)
			}

			for _, repo := range repos.Repos {
				wg.Add(1)
				go func(id int, name string) {
					defer wg.Done()
					issues := &Issues{}
					issues.OrgId = org.Id
					issues.RepoId = id
					err = update(c.FetchAll(fmt.Sprintf(issuesUrl, org.Login, name), &issues.Issues), b, issues)
					if err != nil {
						log.Printf("error fetching issues for <%s>: %v", name, err)
						return
					}
					if err = b.Insert(issues); err != nil {
						log.Printf("error updating issues for <%s>: %v", name, err)
					}
				}(repo.Id, repo.Name)
			}

		}
	}()

	wg.Wait()

	return nil
}
