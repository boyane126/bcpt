package main

import (
	"log"
	"os"

	"github.com/boyane126/bcpt/internal/bcptctl"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	command := bcptctl.NewDefaultBCPTCommand()
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
