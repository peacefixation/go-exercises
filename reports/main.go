package main

import (
	"exercises/reports/database"
	"exercises/reports/report"
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -report reportName\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	reportName := flag.String("report", "", "the name of the report to run")
	flag.Parse()

	if *reportName == "" {
		usage()
	}

	r, err := report.Create(*reportName)
	if err != nil {
		log.Fatal(err.Error())
	}

	content, err := r.Run(database.SampleCredentials())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(*content)

	// TODO: send email
}
