package main

import (
	"os"

	"github.com/rarimo/noaa-indexer/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
