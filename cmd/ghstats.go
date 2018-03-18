package main

import (
	"encoding/csv"
	g "github.com/vrde/ghstats"
	"log"
	"os"
)

func main() {
	ctx := g.Context{}
	if token, defined := os.LookupEnv("GITHUB_TOKEN"); defined {
		ctx.GitHubToken = token
	} else {
		log.Fatal("GITHUB_TOKEN env variable not found.")
	}

	ch := make(chan *g.IssuesResponse)
	go g.FetchIssues(&ctx, os.Args[1], ch)

	w := csv.NewWriter(os.Stdout)
	w.Write(g.IssueHeaders)
	for i := range ch {
		if i.Error != nil {
			log.Printf("Error retrieving <%s>: %v", i.Url, i.Error)
		} else {
			w.WriteAll(i.Issues.ToSlice())
		}
	}
}
