package main

import (
	"log"

	"github.com/o-richard/fit/pkg/db"
	"github.com/o-richard/fit/pkg/server"
)

func main() {
	appdb, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	server.StartServer("8000")
	appdb.Close()
}
