package shared

import (
	"log"
	"os"
)

var L = log.New(os.Stdout, "", log.LstdFlags)
