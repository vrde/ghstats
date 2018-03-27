package main

import (
	g "github.com/vrde/ghstats"
	"log"
	"os"
)

func main() {
	ctx := g.GetContext()
	backend := g.GetBackend()
	err := g.UpdateAllFromOrg(ctx, backend, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
