package main

import (
	"encoding/csv"
	g "github.com/vrde/ghstats"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)
	ch := make(chan *g.Issues)
	h := true

	go g.FetchIssues(os.Args[1], ch)

	for issues := range ch {
		if h {
			w.Write(issues.GetHeaders())
			h = false
		}
		w.WriteAll(issues.ToSlice())
	}
}
