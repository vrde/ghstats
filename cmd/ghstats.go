package main

import (
	"encoding/csv"
	"github.com/vrde/gitstats"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)
	ch := make(chan *gitstats.Issues)
	h := true
	go gitstats.FetchIssues(os.Args[1], ch)
	for issues := range ch {
		if h {
			w.Write(issues.GetHeaders())
			h = false
		}
		w.WriteAll(issues.ToSlice())
	}
}
