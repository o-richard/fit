package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/o-richard/fit/pkg/db"
	"github.com/o-richard/fit/pkg/parser"
	"github.com/o-richard/fit/pkg/server"
)

func main() {
	appdb, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer appdb.Close()

	var fitnessAppParser, serverPort string
	flag.StringVar(&fitnessAppParser, "parse", "", "Name of the parser. Choices are samsung.")
	flag.StringVar(&serverPort, "server", "8000", "Start the HTTP server at the provided port.")
	flag.Parse()
	if maxArgs := 2; len(os.Args) < maxArgs {
		flag.Usage()
		return
	}

	if fitnessAppParser != "" {
		err := parser.ParseFitnessAppRecords(fitnessAppParser)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("The fitness app records in the database are up to date.")
	}
	if serverPort != "" {
		if _, err := strconv.Atoi(serverPort); err != nil {
			fmt.Println("please provide a valid port number")
			return
		}
		server.StartServer(serverPort)
	}
}
