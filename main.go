package main

import (
	"log"
	"os"

	"github.com/enginyoyen/mimic/pkg/cmd"
	"github.com/enginyoyen/mimic/pkg/mapping"
	"github.com/enginyoyen/mimic/pkg/server"
)

func main() {
	config, app := cmd.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	files, err := mapping.WalkFiles(config.MappingDirectory)
	if err != nil {
		log.Fatal(err)
	}

	mappings := mapping.MapRequests(files)

	server.Serve(config, mappings)
}
