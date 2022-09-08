package main

import (
	"flag"
	"log"
	"os"

	"github.com/HonbraDev/soical/generator"
)

var (
	outputFile = flag.String("o", "calendar.ics", "output file")
	username   = flag.String("u", "", "username")
	password   = flag.String("p", "", "password")
)

func main() {
	flag.Parse()

	if *username == "" || *password == "" {
		flag.Usage()
		return
	}

	log.Println("Generating calendar")
	cal, err := generator.MakeCalendarLazy(*username, *password)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	f, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("Writing calendar to file")
	if err := cal.SerializeTo(f); err != nil {
		log.Fatal("Error: ", err)
	}
}
