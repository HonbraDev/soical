package main

import (
	"log"

	"github.com/HonbraDev/soical/server"
)

func main() {
	log.Fatal(server.NewServer(":3000").ListenAndServe())
}
