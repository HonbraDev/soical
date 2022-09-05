package main

import (
	"github.com/HonbraDev/soical/server"
	"github.com/HonbraDev/soical/shared"
)

func main() {
	shared.L.Fatal(server.NewServer(":3000").ListenAndServe())
}
