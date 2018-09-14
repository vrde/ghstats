package ghstats

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"text/template"
)

type Stack struct {
	Item interface{}
	Up   *Stack
}

type UrlTemplate struct {
	Name      string
	Template  string
	_template *template.Template
}

func (u *UrlTemplate) Execute(s *Stack) string {
	if u._template == nil {
		var err error
		if u._template, err = template.New(u.Name).Parse(u.Template); err != nil {
			panic(err)
		}
	}
	b := new(bytes.Buffer)
	if err := u._template.Execute(b, s); err != nil {
		panic(err)
	}
	return b.String()
}

type APITree struct {
	UrlTemplate UrlTemplate
	Storer      Storer
	APITree     []APITree
}

var tree = APITree{

	UrlTemplate: UrlTemplate{Name: "orgUrl", Template: "/orgs/{{.Item.Login}}"},
	Storer:      &Org{},
	APITree: []APITree{
		APITree{
			UrlTemplate: UrlTemplate{Name: "membersUrl", Template: "/orgs/{{.Item.Login}}/members"},
			Storer:      &Members{},
		},
		APITree{
			UrlTemplate: UrlTemplate{Name: "reposUrl", Template: "/orgs/{{.Item.Login}}/repos"},
			Storer:      &Repos{},
			APITree: []APITree{
				APITree{
					UrlTemplate: UrlTemplate{Name: "issuesUrl", Template: "/repos/{{.Up.Item.Login}}/{{.Item.Name}}/issues?state=all"},
					Storer:      &Issues{},
				},
			},
		},
	},
}

func FetchOrg(api *API, name string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go recursiveFetch(api, wg, &tree, &Stack{Item: &OrgBuffer{Login: name}, Up: nil})
	wg.Wait()
}

func recursiveFetch(api *API, wg *sync.WaitGroup, t *APITree, s *Stack) {
	defer wg.Done()
	log.Println("Current APITree → ", t)
	out := ""

	for last := s; last != nil; last = last.Up {
		out += fmt.Sprintf("%T, ", last.Item)
	}
	log.Println("Stack → " + out)

	url := api.GitHubRootAPI + t.UrlTemplate.Execute(s)

	for url != "" {
		buffer := t.Storer.NewBuffer()
		next, err := api.Fetch(url, &buffer)
		url = next

		if err != nil {
			panic(err)
		}

		// store in the backend...

		for _, t := range t.APITree {
			for i, item := range buffer.Items() {
				log.Printf("Item (%d) → %T\n", i, item)
				stack := &Stack{Item: item, Up: s}
				wg.Add(1)
				go recursiveFetch(api, wg, &t, stack)
			}
		}
	}
}
