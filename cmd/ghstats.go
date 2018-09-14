package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	g "github.com/vrde/ghstats"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "./ghstats.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api := g.GetAPI()
	backend := g.GetBackend(db)
	if err = backend.CreateTables(&g.Org{}, &g.Members{}, &g.Repos{}, &g.Issues{}); err != nil {
		log.Fatal(err)
	}

	g.FetchOrg(api, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
