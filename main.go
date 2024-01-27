package main

import (
	"log"
	"os"

	"github.com/travisyeah/tess/tess"
)

func main() {
	log.SetFlags(0)

	args := os.Args[1:]
	tess := &tess.Tess{}
	tess.Run(args)
}
