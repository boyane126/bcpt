package main

import (
	"log"
	"os"

	"github.com/boyane126/bcpt/internal/bcptctl/cmd"
)

func main() {
	command := cmd.NewDefaultBCPTCommand()
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
