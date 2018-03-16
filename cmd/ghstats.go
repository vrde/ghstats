package main

import (
	"encoding/csv"
	g "github.com/vrde/ghstats"
	"log"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)
	ch := make(chan *g.Issues)
	h := true
	ctx := g.Context{}
	if token, defined := os.LookupEnv("GITHUB_TOKEN"); defined {
		ctx.GitHubToken = token
	} else {
		log.Fatal("GITHUB_TOKEN env variable not found.")
	}

	go g.FetchIssues(&ctx, os.Args[1], ch)

	for issues := range ch {
		if h {
			w.Write(issues.GetHeaders())
			h = false
		}
		w.WriteAll(issues.ToSlice())
	}
}
