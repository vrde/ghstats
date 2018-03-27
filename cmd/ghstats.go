package main

import (
	g "github.com/vrde/ghstats"
	"log"
	"os"
)

func main() {
	api := g.GetContext()
	backend := g.GetBackend()
	err := g.UpdateAllFromOrg(api, backend, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
