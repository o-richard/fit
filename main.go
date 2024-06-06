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
	parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)
	serverCmd := flag.NewFlagSet("runserver", flag.ExitOnError)

	var parserType string
	parseCmd.StringVar(&parserType, "type", "", "Name of the parser. Choices are samsung.")

	var serverPort string
	serverCmd.StringVar(&serverPort, "port", "8000", "Start the HTTP server at the provided port.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s [command] [options]\n\nAvailable commands:\nparse\n", os.Args[0])
		parseCmd.PrintDefaults()
		fmt.Printf("runserver\n")
		serverCmd.PrintDefaults()
	}

	if maxArgs := 2; len(os.Args) < maxArgs {
		flag.Usage()
		os.Exit(1)
	}

	appdb, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer appdb.Close()

	switch os.Args[1] {
	case "parse":
		_ = parseCmd.Parse(os.Args[2:])
		if err := parser.ParseFitnessAppRecords(parserType); err != nil {
			appdb.Close()
			log.Fatal(err)
		}
		fmt.Println("The fitness app records in the database are up to date.")
	case "runserver":
		_ = serverCmd.Parse(os.Args[2:])
		if _, err := strconv.Atoi(serverPort); err != nil {
			appdb.Close()
			log.Fatal("please provide a valid port number")
		}
		server.StartServer(serverPort)
	default:
		appdb.Close()
		fmt.Println("Unknown command:", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}
